package generate

import (
	_ "embed"
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/api/spec"
	"github.com/bhmy-shm/tools/gofkctl/common/env"
	"github.com/bhmy-shm/tools/gofkctl/common/format"
	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/bhmy-shm/tools/gofkctl/config"
	"strings"
)

//go:embed api_main.tpl
var mainTemplate string

func genMain(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	name := strings.ToLower(api.Service.Name)
	filename, err := format.FileNamingFormat(cfg.NamingFormat, name)
	if err != nil {
		return err
	}

	configName := filename
	if strings.HasSuffix(filename, "-apis") {
		filename = strings.ReplaceAll(filename, "-apis", "")
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        filename + ".go",
		templateName:    "mainTemplate",
		category:        category,
		templateFile:    mainTemplateFile,
		builtinTemplate: mainTemplate,
		data: map[string]string{
			"importPackages": genMainImports(rootPkg, filename),
			"serviceName":    configName,
			"mount":          genMountControl(name),
		},
	})
}

func genMainImports(parentPkg, filename string) string {
	var imports []string
	/*
		"github.com/bhmy-shm/gofks"
		gofkConf "github.com/bhmy-shm/gofks/core/config"

		"test/controls/test"
		"test/wire"
	*/
	imports = append(imports, fmt.Sprintf("\"%s\"", env.ProjectOpenSourceURL))
	imports = append(imports, fmt.Sprintf("gofkConf \"%s/core/config\"", env.ProjectOpenSourceURL))

	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, wireDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, controls+"/"+filename)))

	return strings.Join(imports, "\n\t")
}

func genMountControl(name string) string {
	return fmt.Sprintf("%s.New%sController()", name, strings.Title(name))
}
