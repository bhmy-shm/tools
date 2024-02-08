package rpc

import (
	"fmt"
	"github.com/bhmy-shm/tools/gofkctl/rpc/rpcFlag"
	"github.com/spf13/cobra"
	"log"
)

var f *rpcFlag.RpcProtoFlag

func init() {
	f = rpcFlag.NewRpcProtoFlag()
}

func RpcCmd() *cobra.Command {
	rpcRoot := &cobra.Command{
		Use:          "rpc",
		Short:        "rpc",
		Long:         "rpc [option(new|rpcFlag)]",
		Example:      "gofkctl rpc [option] [flags]",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("gofkctl rpc [option] [flags]", args)
		},
	}

	rpcRoot.AddCommand(protoCmd())
	rpcRoot.AddCommand(rpcNewCmd())

	return rpcRoot
}

func protoCmd() *cobra.Command {

	proto := &cobra.Command{
		Use:          "rpcFlag",
		Short:        "rpcFlag",                              //短帮助信息
		Long:         "rpc rpcFlag [option]",                 //长帮助信息
		Example:      "gofkctl rpc rpcFlag [option] [flags]", //示例
		SilenceUsage: true,
		Run:          protoRun,
	}

	proto.Flags().StringVarP(&f.Source, "source", "s", "./*.rpcFlag", "指定proto原文件")
	proto.Flags().StringVarP(&f.Dest, "dest", "d", ".", "")
	proto.Flags().StringVarP(&f.Types, "types", "t", f.Types, "")
	return proto
}

func protoRun(cmd *cobra.Command, args []string) {

	//校验传入的参数是否定义过
	for _, arg := range args {
		flag := cmd.Flags().Lookup(arg)
		if flag == nil {
			cmd.Printf("%s is defined as a rpcFlag, please -help", arg)
		}
	}

	var (
		name string
	)

	if len(f.Source) > 0 {
		fmt.Println("source arg =", f.Source, f.Source[:len(f.Source)-6])
		//处理 .\group.proto 开头的 .\
		//处理 ./group.proto 开头的 ./

		name = f.Source[:len(f.Source)-6] + "/"
		fmt.Println("filename=", name)
	}

	if len(f.Dest) > 0 {
		fmt.Println("dest arg =", f.Dest)
	}

	if len(f.Types) > 0 {
		fmt.Println("types arg =", f.Types)
	}

	err := buildDirectory(f.Source, name)
	if err != nil {
		log.Printf("build Directory and file is Failed: %s", err)
	}
}
