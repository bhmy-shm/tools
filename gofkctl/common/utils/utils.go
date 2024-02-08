package utils

import (
	"errors"
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/api/spec"
	goformat "go/format"
	"io"
	"os"
	"regexp"
	"strings"
)

func FormatCode(code string) string {
	ret, err := goformat.Source([]byte(code))
	if err != nil {
		return code
	}

	return string(ret)
}

// IsTemplateVariable 函数会返回 true，如果文本是一个模板变量
// 文本必须以点号开头，并且是一个有效的模板。
func IsTemplateVariable(text string) bool {
	match, _ := regexp.MatchString(`(?m)^{{(\.\w+)+}}$`, text)
	return match
}

// TemplateVariable 函数返回模板的变量名。
func TemplateVariable(text string) string {
	if IsTemplateVariable(text) {
		return text[3 : len(text)-2]
	}
	return ""
}

// Copy calls io.Copy if the source file and destination file exists
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// WrapErr wraps an error with message
func WrapErr(err error, message string) error {
	return errors.New(message + ", " + err.Error())
}

// ComponentName returns component name for typescript
func ComponentName(api *spec.ApiSpec) string {
	name := api.Service.Name
	if strings.HasSuffix(name, "-api") {
		return name[:len(name)-4] + "Components"
	}
	return name + "Components"
}

// WriteIndent writes tab spaces
func WriteIndent(writer io.Writer, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Fprint(writer, "\t")
	}
}

// RemoveComment filters comment content
func RemoveComment(line string) string {
	commentIdx := strings.Index(line, "//")
	if commentIdx >= 0 {
		return strings.TrimSpace(line[:commentIdx])
	}
	return strings.TrimSpace(line)
}
