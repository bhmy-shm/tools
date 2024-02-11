package resource

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	gNet "github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"github.com/spf13/cobra"
	"log"
	"net"
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

			data = append(data, genCpuPercent(), genMemory())

			for _, v := range data {
				table.Append(v)
			}

			table.Render()
		}
	}
}

func genCpuPercent() []string {
	p, _ := cpu.Percent(time.Second, VarBoolPercent)
	pCount, _ := cpu.Counts(VarBoolPercent)
	return []string{"CPU", fmt.Sprintf("%d核", pCount), fmt.Sprintf("%.1f%%", p[0])}
}

func genCpuInfo() []string {
	var result []string

	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed,err:%v", err)
	}

	for _, v := range cpuInfos {
		fmt.Printf(v.String())
		result = append(result, v.String())
	}

	//cpu 负载情况
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
	result = append(result, info.String())

	//cpu 核心数
	cpu.Counts(true)  //逻辑核心
	cpu.Counts(false) //物理核心
	return result
}

func genMemory() []string {
	m, _ := mem.VirtualMemory()
	return []string{"内存", fmt.Sprintf("%dG", m.Total/1024/1024/1024), fmt.Sprintf("%.1f%%", m.UsedPercent)}
}

//func genHost() []string {
//	hInfo, _ := host.Info()
//
//	return []string{"host",
//		fmt.Sprintf("操作系统：%v", hInfo.Platform),
//		fmt.Sprintf("info:%v", hInfo),
//		fmt.Sprintf("updateTime:%v, boottime:%v\n", hInfo.Uptime, hInfo.BootTime)}
//}

func genNetIO() []string {
	var result []string
	info, _ := gNet.IOCounters(true)
	for index, v := range info {
		result = append(result, fmt.Sprintf("%v:%v send:%v, recv:%v", index, v, v.BytesRecv, v.BytesRecv))
	}
	return result
}

func genLocalIp() []string {
	var result []string
	addrSlice, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, addr := range addrSlice {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		result = append(result, ipAddr.IP.String())
	}
	return result
}

func genNetBoundIp() []string {
	var result []string
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	result = append(result, localAddr.String())

	return result
}

func genDisk() []string {

	var result []string

	//所有分区
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get partitions failed:err:%v", err)
		return nil
	}

	for _, part := range parts {
		//指定某路径的硬盘使用情况
		diskInfo, _ := disk.Usage(part.Mountpoint)
		result = append(result, fmt.Sprintf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free))
	}

	//所有磁盘IO信息
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		result = append(result, fmt.Sprintf("%v:%v", k, v))
	}

	return result
}

func genPid() {

	pidSlice, _ := process.Pids() //获取所有进程的pids

	for _, pid := range pidSlice {
		pidInfo, _ := process.NewProcess(pid) //指定进程信息
		parent, _ := pidInfo.Parent()         //父进程信息

		log.Println(pidInfo, parent.Pid)
	}
}

func getMemInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 总内存(GB)
	memTotal := memInfo.Total / 1024 / 1024 / 1024
	fmt.Println(fmt.Sprintf("总内存: %v GB", memTotal))

	// 已用内存(MB)
	memUsed := memInfo.Used / 1024 / 1024
	fmt.Println(fmt.Sprintf("已用内存: %v MB", memUsed))

	// 可用内存(MB)
	memAva := memInfo.Available / 1024 / 1024
	fmt.Println(fmt.Sprintf("可用内存: %v MB", memAva))

	// 内存使用率
	memUsedPercent := memInfo.UsedPercent
	fmt.Println(fmt.Sprintf("内存使用率: %.1f%%", memUsedPercent))
}
