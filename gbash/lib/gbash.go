package gbash

import (
	"github.com/progrium/go-basher"
	"log"
	"sync"
)

var (
	BashCtx  *basher.Context
	BashOnce *sync.Once
)

const (
	SHPath   = "/bin/sh"
	BASHPath = "/bin/bash"
)

func NewBash(shellPath string) *basher.Context {

	BashOnce.Do(func() {
		bash, err := basher.NewContext(SHPath, true)
		if err != nil {
			log.Fatal(err)
		}

		BashCtx = bash
	})

	return BashCtx
}

func AddExport() {}

func AddExportFunc() {}
