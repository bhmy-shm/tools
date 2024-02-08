package generate

import (
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
)

const (
	category             = "api"
	configTemplateFile   = "application.tpl"
	contextTemplateFile  = "context.tpl"
	controlsTemplateFile = "controls.tpl"
	typesTemplateFile    = "types.tpl"
	wireTemplateFile     = "wire.tpl"
	mainTemplateFile     = "main.tpl"
	middleTemplateFile   = "middle.tpl"
)

var templates = map[string]string{
	configTemplateFile:   "application.tpl",
	contextTemplateFile:  "context.tpl",
	controlsTemplateFile: "controls.tpl",
	typesTemplateFile:    "types.tpl",
	wireTemplateFile:     "wire.tpl",
	mainTemplateFile:     "main.tpl",
}

// Category returns the category of the api files.
func Category() string {
	return category
}

// Clean cleans the generated deployment files.
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates generates api template files.
func GenTemplates() error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the given template file to the default value.
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}
	return pathx.CreateTemplate(category, name, content)
}

// Update updates the template files to the templates built in current goctl.
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}
