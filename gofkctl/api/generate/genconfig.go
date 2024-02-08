package generate

import (
	_ "embed"
	"github.com/bhmy-shm/tools/gofkctl/common/format"
	"github.com/bhmy-shm/tools/gofkctl/config"
)

const (
	configFile = "application"
)

//go:embed api_application.tpl
var configTemplate string

func genConfig(dir string, cfg *config.Config) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, configFile)
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        filename + ".yaml",
		templateName:    "configTemplate",
		category:        category,
		templateFile:    configTemplateFile,
		builtinTemplate: configTemplate,
		//data: map[string]string{
		//	"authImport": authImportStr,
		//	"auth":       strings.Join(auths, "\n"),
		//	"jwtTrans":   strings.Join(jwtTransList, "\n"),
		//},
	})
}
