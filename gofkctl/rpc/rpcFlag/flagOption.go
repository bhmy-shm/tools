package rpcFlag

type RpcProtoFlag struct {
	Source string //原文件路径
	Dest   string //目标文件路径
	Types  string //service or types , all
}

func NewRpcProtoFlag() *RpcProtoFlag {
	return &RpcProtoFlag{
		Types: "all",
	}
}
