package resource

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/shirou/gopsutil/process"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	VarInt32Process int32
)

//获取指定进程的CPU占用率

func CpuProcess(_ *cobra.Command, args []string) error {

	if VarInt32Process == 0 {
		VarInt32Process = int32(os.Getpid())
	}

	//获取进程信息
	p, _ := process.NewProcess(VarInt32Process)

	//进程名称
	name, _ := p.Name()

	cpuTime, _ := p.Times()
	spew.Dump("p.Times()", cpuTime)

	cpuP, _ := p.CPUPercent()
	fmt.Printf("[%s]进程的cpu占用率:%v, 时间:%v\n", name, cpuP, time.Now().Format(time.DateTime))

	return nil
}
