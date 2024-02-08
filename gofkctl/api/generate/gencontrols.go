package generate

import (
	_ "embed"
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/api/spec"
	"github.com/bhmy-shm/tools/gofkctl/common/env"
	"github.com/bhmy-shm/tools/gofkctl/common/format"
	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/bhmy-shm/tools/gofkctl/common/utils"
	"github.com/bhmy-shm/tools/gofkctl/config"
	"strings"
)

const controlsFileName = "controls"

//go:embed api_controls.tpl
var controlsTemplate string

func genControls(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	goFile, err := format.FileNamingFormat(cfg.NamingFormat, controlsFileName)
	if err != nil {
		return err
	}

	control := strings.Title(api.Service.Name)

	subDir := utils.ToLower(api.Service.Name)

	imports := genControlImports(rootPkg)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          goFile + "/" + subDir,
		filename:        goFile + ".go",
		templateName:    "logicTemplate",
		category:        category,
		templateFile:    controlsTemplateFile,
		builtinTemplate: controlsTemplate,
		data: map[string]string{
			"pkgName": subDir,
			"imports": imports,
			"control": control,
			"handler": genHandlerBuild(api),
		},
	})
}

func genHandlerBuild(api *spec.ApiSpec) string {

	var (
		result      []string
		rootName    = api.Service.Name
		serviceName = api.Service.Name
	)

	result = append(result, fmt.Sprintf("%s:= gofk.Group(\"%s\")", serviceName, utils.ToLower(serviceName)))

	for _, group := range api.Service.Groups {

		gHeader := utils.ToLastLower(group.GetAnnotation("group"))
		if len(gHeader) > 0 {
			serviceName = gHeader
			result = append(result, fmt.Sprintf("\n%s:= %s.Group(\"%s\")", serviceName, rootName, utils.ToLower(gHeader)))
		}
		for _, router := range group.Routes {
			handler := fmt.Sprintf("%s.Handle(\"%s\",\"%s\",s.%s)", serviceName, utils.ToUpper(router.Method), router.Path, router.Handler)
			result = append(result, handler)
		}
	}
	return strings.Join(result, "\n")
}

func genControlImports(parentPkg string) string {
	var imports []string
	/*
		"github.com/bhmy-shm/gofks/example/api/wire"
		"github.com/bhmy-shm/gofks/gofks"
	*/
	imports = append(imports, fmt.Sprintf("\"%s\"", env.ProjectOpenSourceURL))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, wireDir)))
	//imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(parentPkg, typesDir)))
	return strings.Join(imports, "\n\t")
}
