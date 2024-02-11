package common

import (
	"io"
	"os"
	"regexp"
	"strings"
)

// ShellLoader 加载指定路径的脚本文件，并去除脚本文件开头的解释器指定行
func ShellLoader(f string) ([]byte, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`#!.*\n`)
	b = re.ReplaceAll(b, []byte(""))

	return b, nil
}

// ReadStdin 读取 stdin 的内容
func ReadStdin() string {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		return ""
	}
	in := strings.Trim(string(bytes), "\n")
	in = strings.Trim(in, " ")
	return in
}
