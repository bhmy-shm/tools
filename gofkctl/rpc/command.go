package rpc

import (
	"github.com/bhmy-shm/tools/gofkctl/common"
)

func ProtocCommand(protoName string) error {
	return common.ExecuteCommand("protoc", "--go_out=.", protoName)
}

func ProtoGrpcCommand(proto, grpc string) error {
	return common.ExecuteCommand("protoc", "--go_out=.", proto, "--grpc_out=.", grpc)
}

func IsEmpty(s string) bool {
	return len(s) > 0
}
