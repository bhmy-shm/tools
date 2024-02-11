package resource

import (
	"github.com/bhmy-shm/tools/gtop/common"
)

var (
	Cmd    = common.NewCommand("resource", common.WithRunE(ServiceCommand))
	CpuCmd = common.NewCommand("cpu", common.WithRunE(ServiceCommand))
)

func init() {
	var (
		cpuCmdFlags = CpuCmd.Flags()
	)

	cpuCmdFlags.BoolVar(&VarBoolPercent, "percent")

	Cmd.AddCommand(CpuCmd)
}
