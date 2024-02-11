package resource

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	VarBoolPercent bool //false 单核cpu true 多核cpu
)

func ServiceCommand(_ *cobra.Command, args []string) error {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"项目", "数量", "百分比"})

	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			table.ClearRows()

			data := [][]string{}

			data = append(data, genCpu(), genMemory())

			for _, v := range data {
				table.Append(v)
			}

			table.Render()
		}
	}
}

func genCpu() []string {
	p, _ := cpu.Percent(time.Second, VarBoolPercent)
	pCount, _ := cpu.Counts(VarBoolPercent)
	return []string{"CPU", fmt.Sprintf("%d核", pCount), fmt.Sprintf("%.1f%%", p[0])}
}

func genMemory() []string {
	m, _ := mem.VirtualMemory()
	return []string{"内存", fmt.Sprintf("%dG", m.Total/1024/1024/1024), fmt.Sprintf("%.1f%%", m.UsedPercent)}
}
