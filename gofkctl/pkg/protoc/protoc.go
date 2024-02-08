package protoc

import (
	"archive/zip"
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/common/env"
	"github.com/bhmy-shm/tools/gofkctl/common/zipx"
	"github.com/bhmy-shm/tools/gofkctl/pkg/downloader"
	"github.com/bhmy-shm/tools/gofkctl/pkg/goctl"
	"github.com/bhmy-shm/tools/gofkctl/rpc/execx"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var url = map[string]string{
	"linux_32":   "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_32.zip",
	"linux_64":   "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip",
	"darwin":     "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-osx-x86_64.zip",
	"windows_32": "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-win32.zip",
	"windows_64": "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-win64.zip",
}

const (
	Name        = "protoc"
	ZipFileName = Name + ".zip"
)

func Install(cacheDir string) (string, error) {
	return goctl.Install(cacheDir, Name, func(dest string) (string, error) {
		goos := runtime.GOOS
		tempFile := filepath.Join(os.TempDir(), ZipFileName)
		bit := 32 << (^uint(0) >> 63)
		var downloadUrl string
		switch goos {
		case env.OsMac:
			downloadUrl = url[env.OsMac]
		case env.OsWindows:
			downloadUrl = url[fmt.Sprintf("%s_%d", env.OsWindows, bit)]
		case env.OsLinux:
			downloadUrl = url[fmt.Sprintf("%s_%d", env.OsLinux, bit)]
		default:
			return "", fmt.Errorf("unsupport OS: %q", goos)
		}

		err := downloader.Download(downloadUrl, tempFile)
		if err != nil {
			return "", err
		}

		return dest, zipx.Unpacking(tempFile, filepath.Dir(dest), func(f *zip.File) bool {
			return filepath.Base(f.Name) == filepath.Base(dest)
		})
	})
}

func Exists() bool {
	_, err := env.LookUpProtoc()
	return err == nil
}

func Version() (string, error) {
	path, err := env.LookUpProtoc()
	if err != nil {
		return "", err
	}
	version, err := execx.Run(path+" --version", "")
	if err != nil {
		return "", err
	}
	fields := strings.Fields(version)
	if len(fields) > 1 {
		return fields[1], nil
	}
	return "", nil
}
