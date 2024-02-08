package api

import (
	"github.com/bhmy-shm/tools/gofkctl/api/generate"
	apiNew "github.com/bhmy-shm/tools/gofkctl/api/new"
	"github.com/bhmy-shm/tools/gofkctl/common"
	"github.com/bhmy-shm/tools/gofkctl/config"
)

var (
	Cmd       = common.NewCommand("api", common.WithRunE(nil))
	newCmd    = common.NewCommand("new", common.WithRunE(apiNew.CreateServiceCommand))
	generaCmd = common.NewCommand("generate", common.WithRunE(generate.GoCommand))
)

func init() {

	var (
		//apiCmdFlags = Cmd.Flags()
		newCmdFlags    = newCmd.Flags()
		generaCmdFlage = generaCmd.Flags()
	)

	newCmdFlags.StringVar(&apiNew.VarStringHome, "home")
	newCmdFlags.StringVar(&apiNew.VarStringRemote, "remote")
	newCmdFlags.StringVar(&apiNew.VarStringBranch, "branch")
	newCmdFlags.StringVarWithDefaultValue(&apiNew.VarStringStyle, "style", common.DefaultFormat)

	generaCmdFlage.StringVar(&generate.VarStringDir, "dir")
	generaCmdFlage.StringVar(&generate.VarStringAPI, "api")
	generaCmdFlage.StringVar(&generate.VarStringHome, "home")
	generaCmdFlage.StringVar(&generate.VarStringRemote, "remote")
	generaCmdFlage.StringVar(&generate.VarStringBranch, "branch")
	generaCmdFlage.StringVarWithDefaultValue(&generate.VarStringStyle, "style", config.DefaultFormat)

	Cmd.AddCommand(newCmd, generaCmd)
}
