package rpc

import (
	"fmt"
)

type DirectType string

func AddDirect(s string) DirectType {
	return DirectType(s)
}

func (d DirectType) path() string {
	return string(d)
}

type Directory struct {
	directory []DirectType
}

func NewDirectory(pathName string) *Directory {
	return &Directory{
		directory: []DirectType{
			AddDirect("./protoJson/" + pathName),
			AddDirect("./internal/services/" + pathName),
			AddDirect("./internal/validators/" + pathName),
			AddDirect("./internal/storage/" + pathName),
		},
	}
}

func (this *Directory) Range() []DirectType {
	return this.directory
}

// @request option.Source = .rpcFlag 文件
// @request name 是取 .rpcFlag 文件的前缀
func buildDirectory(protoPath, filename string) error {
	var err error
	directories := NewDirectory(filename)

	for _, dir := range directories.Range() {
		fmt.Println("dir path =", dir.path())
		fmt.Println("./internal/services/" + filename)
		fmt.Println("./internal/storage/" + filename)
		fmt.Println("./internal/validators/" + filename)
		// 先生成目录
		err = path.CreateDirectoryIfNotExist(dir.path())
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
		} else {
			fmt.Printf("Directory %s created successfully or already exists.\n", dir)
		}

		//再依次生成目录下文件内容
		switch dir.path() {
		case "./protoJson/" + filename:
			if err = ProtocCommand(protoPath); err != nil {
				return err
			}
		case "./internal/services/" + filename:
			if err = ServicesRpc(dir.path(), filename[:len(filename)-1], protoPath); err != nil {
				return err
			}
		default:
			continue
		}
	}

	return err
}
