package generate

import (
	_ "embed"
	"github.com/bhmy-shm/tools/gofkctl/api/spec"
	"github.com/bhmy-shm/tools/gofkctl/common/format"
	"github.com/bhmy-shm/tools/gofkctl/config"
	"strings"
)

//go:embed api_middleware.tpl
var middlewareImplementCode string

func genMiddleware(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	middlewares := getMiddleware(api)

	for _, item := range middlewares {
		middlewareFilename := strings.TrimSuffix(strings.ToLower(item), "middleware") + "_middleware"
		filename, err := format.FileNamingFormat(cfg.NamingFormat, middlewareFilename)
		if err != nil {
			return err
		}

		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		err = genFile(fileGenConfig{
			dir:             dir,
			subdir:          middlewareDir,
			filename:        filename + ".go",
			templateName:    "contextTemplate",
			category:        category,
			templateFile:    middleTemplateFile,
			builtinTemplate: middlewareImplementCode,
			data: map[string]string{
				"name": strings.Title(name),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
