package vars

import (
	"gbash/lib/common"
	"github.com/progrium/go-basher"
	"github.com/shirou/gopsutil/mem"
	"log"
	"strconv"
)

type MemInfo struct{}

func NewMemInfo() *MemInfo {
	return &MemInfo{}
}

func (m *MemInfo) Info() common.ListExports {
	return func(ctx *basher.Context) {

		v, err := mem.VirtualMemory()
		if err != nil {
			log.Fatal(err)
		}

		ctx.Export("_mem_ava", strconv.FormatUint(v.Available, 10))
		ctx.Export("_mem_total", strconv.FormatUint(v.Total, 10))
		ctx.Export("_mem_used", strconv.FormatUint(v.Used, 10))
	}
}
