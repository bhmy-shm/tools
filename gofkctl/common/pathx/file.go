package pathx

import (
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/common/version"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

var ctlHome string

const (
	NL              = "\n"
	ctlDir          = ".gofkctl"
	gitDir          = ".git"
	cacheDir        = "cache"
	autoCompleteDir = ".auto_complete"
)

var goctlHome string

// RegisterGoctlHome register goctl home path.
func RegisterGoctlHome(home string) {
	goctlHome = home
}

// InitTemplates creates template files GoctlHome where could get it by GetGoctlHome.
func InitTemplates(category string, templates map[string]string) error {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return err
	}

	if err := CreateDirectoryIfNotExist(dir); err != nil {
		return err
	}

	for k, v := range templates {
		if err := createTemplate(filepath.Join(dir, k), v, false); err != nil {
			return err
		}
	}

	return nil
}

// CreateTemplate writes template into file even it is exists.
func CreateTemplate(category, name, content string) error {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return err
	}
	return createTemplate(filepath.Join(dir, name), content, true)
}

// GetDefaultCtlHome returns the path value of the goctl home where Join $HOME with .goctl.
func GetDefaultCtlHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ctlDir), nil
}

// GetGitHome returns the git home of goctl.
func GetGitHome() (string, error) {
	goctlH, err := GetCtlHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(goctlH, gitDir), nil
}

// GetCtlHome returns the path value of the goctl, the default path is ~/.goctl, if the path has
// been set by calling the RegisterGoctlHome method, the user-defined path refers to.
func GetCtlHome() (home string, err error) {
	defer func() {
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		if err == nil && !info.IsDir() {
			os.Rename(home, home+".old")
			CreateDirectoryIfNotExist(home)
		}
	}()
	if len(ctlHome) != 0 {
		home = ctlHome
		return
	}
	home, err = GetDefaultCtlHome()
	return
}

// GetTemplateDir returns the category path value in GoctlHome where could get it by GetctlHome.
func GetTemplateDir(category string) (string, error) {
	home, err := GetCtlHome()
	if err != nil {
		return "", err
	}
	if home == ctlHome {
		// backward compatible, it will be removed in the feature
		// backward compatible start.
		beforeTemplateDir := filepath.Join(home, version.GetVersion(), category)
		entries, _ := os.ReadDir(beforeTemplateDir)
		infos := make([]fs.FileInfo, 0, len(entries))
		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
		var hasContent bool
		for _, e := range infos {
			if e.Size() > 0 {
				hasContent = true
			}
		}
		if hasContent {
			return beforeTemplateDir, nil
		}
		// backward compatible end.

		return filepath.Join(home, category), nil
	}

	return filepath.Join(home, version.GetVersion(), category), nil
}

// LoadTemplate gets template content by the specified file.
func LoadTemplate(category, file, builtin string) (string, error) {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return "", err
	}

	file = filepath.Join(dir, file)
	if !FileExists(file) {
		return builtin, nil
	}

	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Clean deletes all templates and removes the parent directory.
func Clean(category string) error {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

func createTemplate(file, content string, force bool) error {
	if FileExists(file) && !force {
		return nil
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

// MaybeCreateFile creates file if not exists
func MaybeCreateFile(dir, subdir, file string) (fp *os.File, created bool, err error) {

	err = CreateDirectoryIfNotExist(path.Join(dir, subdir))
	if err != nil {
		panic(err)
	}

	fpath := path.Join(dir, subdir, file)
	if FileExists(fpath) {
		fmt.Printf("%s exists, ignored generation\n", fpath)
		return nil, false, nil
	}

	fp, err = CreateFileIfNotExist(fpath)
	created = err == nil
	return
}

// GetCacheDir returns the cache dit of goctl.
func GetCacheDir() (string, error) {
	goctlH, err := GetCtlHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(goctlH, cacheDir), nil
}

func Copy(src, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	dir := filepath.Dir(dest)
	err = CreateDirectoryIfNotExist(dir)
	if err != nil {
		return err
	}
	w, err := os.Create(dest)
	if err != nil {
		return err
	}
	w.Chmod(os.ModePerm)
	defer w.Close()
	_, err = io.Copy(w, f)
	return err
}

// SameFile compares the between path if the same path,
// it maybe the same path in case case-ignore, such as:
// /Users/go_zero and /Users/Go_zero, as far as we know,
// this case maybe appear on macOS and Windows.
func SameFile(path1, path2 string) (bool, error) {
	stat1, err := os.Stat(path1)
	if err != nil {
		return false, err
	}

	stat2, err := os.Stat(path2)
	if err != nil {
		return false, err
	}

	return os.SameFile(stat1, stat2), nil
}
