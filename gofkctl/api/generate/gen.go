package generate

import (
	"errors"
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/api/format"
	"github.com/bhmy-shm/tools/gofkctl/api/parser"
	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/bhmy-shm/tools/gofkctl/common/utils"
	"github.com/bhmy-shm/tools/gofkctl/config"
	"github.com/bhmy-shm/tools/gofkctl/pkg/golang"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const tmpFile = "%s-%d"

var (
	tmpDir = path.Join(os.TempDir(), "gofkctl")
	// VarStringDir describes the directory.
	VarStringDir string
	// VarStringAPI describes the API.
	VarStringAPI string
	// VarStringHome describes the go home.
	VarStringHome string
	// VarStringRemote describes the remote git repository.
	VarStringRemote string
	// VarStringBranch describes the branch.
	VarStringBranch string
	// VarStringStyle describes the style of output files.
	VarStringStyle string
)

// GoCommand gen go project files from command line
func GoCommand(_ *cobra.Command, _ []string) error {
	apiFile := VarStringAPI
	dir := VarStringDir
	namingStyle := VarStringStyle
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	if len(remote) > 0 {
		repo, _ := utils.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}
	if len(apiFile) == 0 {
		return errors.New("missing -api")
	}
	if len(dir) == 0 {
		return errors.New("missing -dir")
	}

	return DoGenProject(apiFile, dir, namingStyle)
}

// DoGenProject gen go project files with api file
func DoGenProject(apiFile, dir, style string) error {

	api, err := parser.Parse(apiFile)
	if err != nil {
		return err
	}

	if err := api.Validate(); err != nil {
		return err
	}

	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	err = pathx.CreateDirectoryIfNotExist(dir)
	if err != nil {
		log.Fatal(err)
	}

	rootPkg, err := golang.GetParentPackage(dir)
	if err != nil {
		return err
	}

	err = genConfig(dir, cfg)
	if err != nil {
		log.Println("genProject config failed:", err)
		return err
	}

	err = genMain(dir, rootPkg, cfg, api)
	if err != nil {
		log.Println("genProject main failed:", err)
		return err
	}

	err = genServiceContext(dir, rootPkg, cfg, api)
	if err != nil {
		log.Println("genProject serviceContext failed:", err)
		return err
	}

	err = genServiceWire(dir, rootPkg, cfg, api)
	if err != nil {
		log.Println("genProject serviceWire failed:", err)
		return err
	}

	err = genTypes(dir, cfg, api)
	if err != nil {
		log.Println("genProject types failed:", err)
		return err
	}

	err = genControls(dir, rootPkg, cfg, api)
	if err != nil {
		log.Println("genProject controls failed:", err)
		return err
	}

	err = genMiddleware(dir, cfg, api)
	if err != nil {
		return err
	}

	if err := backupAndSweep(apiFile); err != nil {
		return err
	}

	if err := format.ApiFormatByPath(apiFile, false); err != nil {
		return err
	}

	fmt.Println(color.Green.Render("Done."))
	return nil
}

func backupAndSweep(apiFile string) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(2)
	_ = os.MkdirAll(tmpDir, os.ModePerm)

	go func() {
		_, fileName := filepath.Split(apiFile)
		_, e := utils.Copy(apiFile, fmt.Sprintf(path.Join(tmpDir, tmpFile), fileName, time.Now().Unix()))
		if e != nil {
			err = e
		}
		wg.Done()
	}()
	go func() {
		if e := sweep(); e != nil {
			err = e
		}
		wg.Done()
	}()
	wg.Wait()

	return err
}

func sweep() error {
	keepTime := time.Now().AddDate(0, 0, -7)
	return filepath.Walk(tmpDir, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		pos := strings.LastIndexByte(info.Name(), '-')
		if pos > 0 {
			timestamp := info.Name()[pos+1:]
			seconds, err := strconv.ParseInt(timestamp, 10, 64)
			if err != nil {
				// print error and ignore
				fmt.Println(color.Red.Sprintf("sweep ignored file: %s", fpath))
				return nil
			}

			tm := time.Unix(seconds, 0)
			if tm.Before(keepTime) {
				if err := os.RemoveAll(fpath); err != nil {
					fmt.Println(color.Red.Sprintf("failed to remove file: %s", fpath))
					return err
				}
			}
		}

		return nil
	})
}
