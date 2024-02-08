package generate

import (
	_ "embed"
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/api/spec"
	"github.com/bhmy-shm/tools/gofkctl/common/env"
	"github.com/bhmy-shm/tools/gofkctl/common/format"
	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/bhmy-shm/tools/gofkctl/config"
)

const wireFileName = "service_wire"

//go:embed api_wire.tpl
var wireTemplate string

func genServiceWire(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, wireFileName)
	if err != nil {
		return err
	}

	configImport := fmt.Sprintf("gofkConf \"%s/core/config\"", env.ProjectOpenSourceURL)
	svcImport := fmt.Sprintf("\"%s\"", pathx.JoinPackages(rootPkg, contextDir))

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          wireDir,
		filename:        filename + ".go",
		templateName:    "wireTemplate",
		category:        category,
		templateFile:    wireTemplateFile,
		builtinTemplate: wireTemplate,
		data: map[string]string{
			"svcImport":    svcImport,
			"configImport": configImport,
			"config":       "gofkConf.Config",
		},
	})
}
