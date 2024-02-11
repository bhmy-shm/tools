package funcs

import (
	"bufio"
	"fmt"
	"gbash/lib/common"
	"github.com/progrium/go-basher"
	"os"
	"strings"
)

type StringFunc struct {
}

func NewString() *StringFunc {
	return &StringFunc{}
}

func (s *StringFunc) Upper() common.ListExports {
	return func(ctx *basher.Context) {
		ctx.ExportFunc("upper", s.toUpper)
	}
}

func (s *StringFunc) toUpper(args []string) {

	//将脚本参数改为大写的
	if len(args) == 1 {
		fmt.Println(strings.ToUpper(args[0]))
	} else {
		//支持管道模式
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			in := ""
			for scanner.Scan() {
				in += scanner.Text()
			}
			if scanner.Err() == nil && in != "" {
				fmt.Println(strings.ToUpper(in))
			}
		}
	}
}
