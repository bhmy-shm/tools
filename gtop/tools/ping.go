package main

import "github.com/bhmy-shm/tools/gtop/common"

// https://github.com/go-ping/ping
var (
	Cmd = common.NewCommand("resource", common.WithRunE(ServiceCommand))
)

func init() {
	var (
		cpuCmdFlags = CpuCmd.Flags()
	)

	cpuCmdFlags.BoolVar(&VarBoolPercent, "percent")

	Cmd.AddCommand(CpuCmd)
}
