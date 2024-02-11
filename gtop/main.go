package main

import (
	"fmt"
	"github.com/bhmy-shm/tools/gtop/common"
	"github.com/bhmy-shm/tools/gtop/resource"
	"github.com/bhmy-shm/tools/gtop/tools"
	"github.com/gookit/color"
	"os"
)

var (
	rootCmd = common.NewCommand("gtop")
)

func init() {
	rootCmd.AddCommand(resource.Cmd, tools.Cmd)
	rootCmd.MustInit()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.Red.Render(err.Error()))
		os.Exit(1)
	}
}
