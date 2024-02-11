package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Upper(args []string) {

	//log.Println("upper args:", args)

	//将脚本参数改为大写的
	if len(args) == 1 {
		fmt.Println(strings.ToUpper(args[0]))
	} else {
		//支持管道模式
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		if len(bytes) > 0 {
			in := strings.Trim(string(bytes), "\n")
			if in != "" {
				fmt.Println(strings.ToUpper(string(bytes)))
			}
		}

	}
}
