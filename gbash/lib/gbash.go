package lib

import (
	"gbash/lib/common"
	"github.com/progrium/go-basher"
	"log"
	"os"
)

type (
	GBash struct {
		sourcePath string
		ctx        *basher.Context
	}
)

const (
	BASHPath = "/bin/bash"
)

func NewBash(filePath, shellPath string) *GBash {

	bash := &GBash{
		sourcePath: filePath,
	}

	ctx, err := basher.NewContext(shellPath, true)
	if err != nil {
		log.Fatal(err)
	}

	bash.ctx = ctx

	return bash
}

func (g *GBash) Used(opts ...common.ListExports) *GBash {
	for _, fn := range opts {
		fn(g.ctx)
	}

	return g
}

func (g *GBash) Run() {

	if g.ctx.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	//加入shell 脚本
	err := g.ctx.Source(g.sourcePath, common.ShellLoader)
	if err != nil {
		log.Fatal("bash ctx Source failed:", err)
	}

	//运行加入的shell脚本
	status, err := g.ctx.Run("", []string{}) //固定 os.Args[1]
	if err != nil {
		log.Fatal("bash ctx Run failed:", err)
	}
	os.Exit(status)
}
