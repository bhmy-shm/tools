package pathx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"text/template"
)

const (
	RpcServiceFieldPrefix = "Service"
)

type RpcMethod struct {
	StructFieldName string
	Name            string
	RequestType     string
	ResponseType    string
}

func (this RpcMethod) FuncStr() string {
	return fmt.Sprintf("func (this *%s) %s(req *rpcFlag.%s, resp *rpcFlag.%s) error {",
		this.StructFieldName, this.Name, this.RequestType, this.ResponseType)
}

type TemplateData struct {
	Content []byte

	PackageName string
	ProtoName   string
	FieldName   string

	ModelPath string
	ModelName string

	RpcFilePath string
	RpcMethods  []RpcMethod
}

func NewTemplate(content []byte) *TemplateData {
	return &TemplateData{
		Content: content,
	}
}

func (this *TemplateData) Parse() (*template.Template, error) {
	if this.Content == nil {
		return nil, errors.New("template parse content is exist")
	}

	tmpl, err := template.New("fileTemplate").Parse(string(this.Content))
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return nil, err
	}
	return tmpl, nil
}

func (this *TemplateData) SetPackage(s string) *TemplateData {
	this.PackageName = s
	return this
}

func (this *TemplateData) SetProtoName(s string) *TemplateData {
	this.ProtoName = s
	return this
}

func (this *TemplateData) SetFieldName(s string) *TemplateData {
	this.FieldName = s
	return this
}

func (this *TemplateData) SetModelPath(s string) *TemplateData {
	this.ModelPath = s
	return this
}

func (this *TemplateData) SetModelName(s string) *TemplateData {
	this.ModelName = s
	return this
}

func (this *TemplateData) SetRpcFilePath(s string) *TemplateData {
	this.RpcFilePath = s
	return this
}

func (this *TemplateData) ExtractRPCMethods() []RpcMethod {

	if this.RpcFilePath == "" {
		log.Println("rpc filed is exist")
		return nil
	}

	protoContent, err := os.ReadFile(this.RpcFilePath)
	if err != nil {
		fmt.Printf("Error reading .rpcFlag file: %v\n", err)
		return nil
	}

	str := `rpc\s+(\w+)\s+\((\w+)\)\s+returns\((\w+)\);`
	rpcRegex := regexp.MustCompile(str)
	matches := rpcRegex.FindAllStringSubmatch(string(protoContent), -1)

	for _, match := range matches {
		//fmt.Printf("rpcFlag rpc match:%v", match)
		if len(match) >= 4 {
			rpcName := match[1]
			requestType := match[2]
			responseType := match[3]

			//fmt.Println("RPC Name:", rpcName)
			//fmt.Println("Request Type:", requestType)
			//fmt.Println("Response Type:", responseType)
			this.RpcMethods = append(this.RpcMethods, RpcMethod{
				StructFieldName: RpcServiceFieldPrefix + this.FieldName,
				Name:            rpcName,
				RequestType:     requestType,
				ResponseType:    responseType,
			})
		}
	}

	return this.RpcMethods
}
