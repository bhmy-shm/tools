package ctx

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/bhmy-shm/tools/gofkctl/rpc/execx"
)

const goModuleWithoutGoFiles = "command-line-arguments"

var errInvalidGoMod = errors.New("invalid go module")

// Module contains the relative data of go module,
// which is the result of the command go list
type Module struct {
	Path      string
	Main      bool
	Dir       string
	GoMod     string
	GoVersion string
}

func (m *Module) validate() error {
	if m.Path == goModuleWithoutGoFiles || m.Dir == "" {
		return errInvalidGoMod
	}
	return nil
}

// projectFromGoMod is used to find the go module and project file path
// the workDir flag specifies which folder we need to detect based on
// only valid for go mod project
func projectFromGoMod(workDir string) (*ProjectContext, error) {
	if len(workDir) == 0 {
		return nil, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return nil, err
	}

	workDir, err := pathx.ReadLink(workDir)
	if err != nil {
		return nil, err
	}

	m, err := getRealModule(workDir, execx.Run)
	if err != nil {
		return nil, err
	}
	if err := m.validate(); err != nil {
		return nil, err
	}

	var ret ProjectContext
	ret.WorkDir = workDir
	ret.Name = filepath.Base(m.Dir)
	dir, err := pathx.ReadLink(m.Dir)
	if err != nil {
		return nil, err
	}

	ret.Dir = dir
	ret.Path = m.Path
	return &ret, nil
}

func getRealModule(workDir string, execRun execx.RunFunc) (*Module, error) {
	data, err := execRun("go list -json -m", workDir)
	if err != nil {
		return nil, err
	}

	modules, err := decodePackages(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	for _, m := range modules {
		realDir, err := pathx.ReadLink(m.Dir)
		if err != nil {
			return nil, fmt.Errorf("failed to read go.mod, dir: %s, error: %w", m.Dir, err)
		}

		if strings.HasPrefix(workDir, realDir) {
			return &m, nil
		}
	}

	return nil, errors.New("no matched module")
}

func decodePackages(reader io.Reader) ([]Module, error) {
	br := bufio.NewReader(reader)
	if _, err := br.ReadSlice('{'); err != nil {
		return nil, err
	}

	if err := br.UnreadByte(); err != nil {
		return nil, err
	}

	var modules []Module
	decoder := json.NewDecoder(br)
	for decoder.More() {
		var m Module
		if err := decoder.Decode(&m); err != nil {
			return nil, fmt.Errorf("invalid module: %v", err)
		}

		modules = append(modules, m)
	}

	return modules, nil
}

// projectFromGoPath is used to find the main module and project file path
// the workDir flag specifies which folder we need to detect based on
// only valid for go mod project
func projectFromGoPath(workDir string) (*ProjectContext, error) {
	if len(workDir) == 0 {
		return nil, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return nil, err
	}

	workDir, err := pathx.ReadLink(workDir)
	if err != nil {
		return nil, err
	}

	buildContext := build.Default
	goPath := buildContext.GOPATH
	goPath, err = pathx.ReadLink(goPath)
	if err != nil {
		return nil, err
	}

	goSrc := filepath.Join(goPath, "src")
	if !pathx.FileExists(goSrc) {
		return nil, errModuleCheck
	}

	wd, err := filepath.Abs(workDir)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(wd, goSrc) {
		return nil, errModuleCheck
	}

	projectName := strings.TrimPrefix(wd, goSrc+string(filepath.Separator))
	return &ProjectContext{
		WorkDir: workDir,
		Name:    projectName,
		Path:    projectName,
		Dir:     filepath.Join(goSrc, projectName),
	}, nil
}
