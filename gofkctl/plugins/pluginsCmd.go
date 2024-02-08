package plugins

import "github.com/spf13/cobra"

func PluginsCmd() *cobra.Command {
	plugins := &cobra.Command{
		Use:          "plugins",
		Short:        "plugins",                          //短帮助信息
		Long:         "plugins [option(new)]",            //长帮助信息
		Example:      "gofkctl plugins [option] [flags]", //示例
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("gofkctl plugins [option] [flags]")
		},
	}
	plugins.AddCommand(pluginsNew())
	return plugins
}

func pluginsNew() *cobra.Command {
	return &cobra.Command{
		Use:          "new",
		Short:        "new",                         //短帮助信息
		Long:         "new [option]",                //长帮助信息
		Example:      "gofkctl plugins new [flags]", //示例
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("gofkctl plugins new [flags]")
		},
	}
}
