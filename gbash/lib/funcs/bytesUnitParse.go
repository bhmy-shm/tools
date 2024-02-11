package funcs

import (
	"fmt"
	"gbash/lib/common"
	"github.com/progrium/go-basher"
	"log"
	"strconv"
	"strings"
)

const (
	UnitK = "K"
	UnitM = "M"
	UnitG = "G"
)

type BytesUnitParse struct{}

func NewBytesUnitParse() *BytesUnitParse {
	return &BytesUnitParse{}
}

func (p *BytesUnitParse) UnitParse() common.ListExports {
	return func(ctx *basher.Context) {
		ctx.ExportFunc(UnitM, func(strings []string) {
			fmt.Println(p.unit(UnitM))
		})
		ctx.ExportFunc(UnitG, func(i []string) {
			fmt.Println(p.unit(UnitG))
		})
		ctx.ExportFunc(UnitK, func(i []string) {
			fmt.Println(p.unit(UnitK))
		})
	}
}

// stdin 当做 byte
func (p *BytesUnitParse) unit(unit string) string {
	none := "0M"
	in := common.ReadStdin()
	if in != "" {

		//转为  uint64
		uin, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			return none
		}

		unit = strings.ToUpper(unit)
		log.Println("unit parse :", unit)
		switch unit {
		case UnitK:
			uin = uin / 1024
		case UnitM:
			uin = uin / 1024 / 1024
		case UnitG:
			uin = uin / 1024 / 1024 / 1024
		default:
			return none
		}
		return strconv.FormatUint(uin, 10) + unit
	}
	return none
}
