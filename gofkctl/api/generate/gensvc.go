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

const contextFileName = "service_context"

//go:embed api_svc.tpl
var contextTemplate string

func genServiceContext(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, contextFileName)
	if err != nil {
		return err
	}

	var middlewareStr string
	var middlewareAssignment string
	middlewares := getMiddleware(api)

	for _, item := range middlewares {
		middlewareStr += fmt.Sprintf("%s rest.Middleware\n", item)
		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		middlewareAssignment += fmt.Sprintf("%s: %s,\n", item,
			fmt.Sprintf("middleware.New%s().%s", strings.Title(name), "Handle"))
	}
	configImport := fmt.Sprintf("gofkConf \"%s/core/config\"", env.ProjectOpenSourceURL)
	if len(middlewareStr) > 0 {
		configImport += "\n\t\"" + pathx.JoinPackages(rootPkg, middlewareDir) + "\""
		configImport += fmt.Sprintf("\n\t\"%s/rest\"", env.ProjectOpenSourceURL)
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          contextDir,
		filename:        filename + ".go",
		templateName:    "contextTemplate",
		category:        category,
		templateFile:    contextTemplateFile,
		builtinTemplate: contextTemplate,
		data: map[string]string{
			"configImport":         configImport,
			"config":               "gofkConf.Config",
			"middleware":           middlewareStr,
			"middlewareAssignment": middlewareAssignment,
		},
	})
}
