package rpc

import "github.com/spf13/cobra"

func rpcNewCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "new",
		Short:        "new",
		Long:         "rpc new",
		Example:      "gofkctl rpc new [flags]",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("gofkctl rpc new [flags]", args)
		},
	}
}
