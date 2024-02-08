package new

import (
	_ "embed"
	"errors"
	"github.com/bhmy-shm/tools/gofkctl/api/generate"
	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed api.tpl
var apiTemplate string

var (
	// VarStringHome gofkctl 的本地目录路径。
	VarStringHome string
	// VarStringRemote 远程 Git 仓库的 URL。
	VarStringRemote string
	// VarStringBranch Git 仓库的分支名称。
	VarStringBranch string
	// VarStringStyle 文件输出的样式
	VarStringStyle string
)

func CreateServiceCommand(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		log.Fatal("Missing api program name")
	}
	//new 的服务名称，如果没有直接报错
	dirName := args[0]
	if strings.Contains(dirName, "-") {
		return errors.New("api new command service name not support strikethrough, because this will used by function name")
	}

	// 获取dirName 的绝对路径
	abs, err := filepath.Abs(dirName)
	if err != nil {
		return err
	}

	//检查绝对路径是否存在，不存在则创建
	err = pathx.CreateDirectoryIfNotExist(abs)
	if err != nil {
		return err
	}

	//拼接绝对路径和文件名以形成 .api 文件的完整路径。并创建 .api 文件
	dirName = filepath.Base(filepath.Clean(abs))
	filename := dirName + ".api"
	apiFilePath := filepath.Join(abs, filename)
	fp, err := os.Create(apiFilePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	//加载模版文件
	text, err := pathx.LoadTemplate(category, apiTemplateFile, apiTemplate)
	if err != nil {
		return err
	}

	//将模版内容输出写入.api文件
	t := template.Must(template.New("template").Parse(text))
	if err := t.Execute(fp, map[string]string{
		"name":    dirName,
		"handler": strings.Title(dirName),
	}); err != nil {
		return err
	}

	//生成项目
	err = generate.DoGenProject(apiFilePath, abs, VarStringStyle)

	return err
}
