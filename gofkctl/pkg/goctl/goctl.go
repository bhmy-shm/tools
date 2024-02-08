package goctl

import (
	"github.com/bhmy-shm/tools/gofkctl/common/env"
	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/bhmy-shm/tools/gofkctl/pkg/golang"
	"log"
	"path/filepath"
	"runtime"
)

func Install(cacheDir, name string, installFn func(dest string) (string, error)) (string, error) {
	goBin := golang.GoBin()
	cacheFile := filepath.Join(cacheDir, name)
	binFile := filepath.Join(goBin, name)

	goos := runtime.GOOS
	if goos == env.OsWindows {
		cacheFile = cacheFile + ".exe"
		binFile = binFile + ".exe"
	}
	// read cache.
	err := pathx.Copy(cacheFile, binFile)
	if err == nil {
		log.Printf("%q installed from cache\n", name)
		return binFile, nil
	}

	binFile, err = installFn(binFile)
	if err != nil {
		return "", err
	}

	// write cache.
	err = pathx.Copy(binFile, cacheFile)
	if err != nil {
		log.Printf("write cache error: %+v\n", err)
	}
	return binFile, nil
}
