package rpc

import (
	"fmt"
	common2 "github.com/bhmy-shm/tools/gofkctl/common"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ServicesRpc(dir, filename, protocName string) error {

	var (
		protoName      = filename
		fieldName      = strings.ToUpper(string(filename[0])) + filename[1:]
		serviceContent = []byte(common2.TemplateRpc[common2.TemplateRpcServer])
	)

	filePath := filepath.Join(dir, filename+".go")

	//先判断要生成的文件，是否已经存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {

		tp := path.NewTemplate(serviceContent).
			SetPackage(filename).
			SetProtoName(protoName).
			SetFieldName(fieldName).
			SetRpcFilePath(protocName)

		tp.ExtractRPCMethods()
		return path.NewFileBuild(tp, filePath)
	}

	//存在，则拿到文件所有内容 []byte
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("读取文件失败，failed: %s", err)
		return err
	}

	//读取rpcMethod模板文件
	rpcMethodsContent := []byte(common2.TemplateRpc[common2.TemplateRpcServerWire]) //todo 这都不对
	tp := path.NewTemplate(rpcMethodsContent).
		SetPackage(filename).
		SetProtoName(protoName).
		SetFieldName(fieldName).
		SetRpcFilePath(protocName)

	content = readTemplateByte(tp, content)

	//此时的 content 就是要生成的文件内容
	if len(content) == 0 {
		return nil
	}

	//传入 content 并生成文件
	return common2.NewServiceWithContent(filePath, content)
}

// --------------------- 处理模板文件生成的覆盖和追加 ----------------------

type MyWriter struct {
	data []byte
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	w.data = append(w.data, p...)
	return len(p), nil
}

func (w *MyWriter) String() string {
	return string(w.data)
}

func readTemplateByte(tp *path.TemplateData, fileBody []byte) []byte {

	//读取proto中的所有方法
	tp.ExtractRPCMethods()

	//循环生成每一个服务方法
	for _, v := range tp.RpcMethods {
		log.Println("rpcMethods v=", v)
		//解析成要加载的rpcMethod方法数据
		tmpl, err := tp.Parse()
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			return nil
		}

		//v当中的字段拼凑出一个函数方法
		rpcFuncStr := v.FuncStr()

		//拿到基于当前方法的内容
		resultContentBuffer := &MyWriter{}
		err = tmpl.Execute(resultContentBuffer, v)
		if err != nil {
			log.Printf("tmpl Execute is failed:%s", err)
			return nil
		}

		log.Println("循环生成的模板方法内容：", resultContentBuffer.String())

		//校验生成内容与拼凑的函数方法能否匹配的上
		if !strings.Contains(resultContentBuffer.String(), rpcFuncStr) {
			log.Println("生成内容与函数方法不一致")
			return nil
		}

		//用当前生成的内容，在所有的文件内容中进行匹配
		if !strings.Contains(string(fileBody), rpcFuncStr) {
			//如果没有则追加到fileBody的末尾处
			fileBody = append(fileBody, '\n')
			fileBody = append(fileBody, resultContentBuffer.data...)
		}
	}

	return fileBody
}
