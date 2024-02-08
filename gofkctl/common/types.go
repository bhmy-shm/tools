package common

type (
	RpcServer string
	ApiServer string
)

const (
	TemplateRpcServer        RpcServer = "TemplateRpcServer"
	TemplateRpcServerClient  RpcServer = "TemplateRpcServerClient"
	TemplateRpcServerLogic   RpcServer = "TemplateRpcServerLogic"
	TemplateRpcServerProtoc  RpcServer = "TemplateRpcServerProtoc"
	TemplateRpcServerContext RpcServer = "TemplateRpcServerContext"
	TemplateRpcServerMain    RpcServer = "TemplateRpcServerMain"
	TemplateRpcServerWire    RpcServer = "TemplateRpcServerWire"
)

const (
	TemplateApiServer        ApiServer = "TemplateApiServer"
	TemplateApiClient        ApiServer = "TemplateApiClient"
	TemplateApiControl       ApiServer = "TemplateApiControl"
	TemplateApiServerContext ApiServer = "TemplateApiServerContext"
	TemplateApiServerMain    ApiServer = "TemplateApiServerMain"
	TemplateApiServerWire    ApiServer = "TemplateApiServerWire"
)

var (
	TemplateRpc map[RpcServer]string
	TemplateApi map[ApiServer]string
)

func init() {
	TemplateRpc = map[RpcServer]string{
		TemplateRpcServer:        "",
		TemplateRpcServerClient:  "",
		TemplateRpcServerLogic:   "",
		TemplateRpcServerProtoc:  "",
		TemplateRpcServerContext: "",
		TemplateRpcServerMain:    "",
		TemplateRpcServerWire:    "",
	}

	TemplateApi = map[ApiServer]string{
		TemplateApiServer:        "",
		TemplateApiClient:        "",
		TemplateApiControl:       "",
		TemplateApiServerContext: "",
		TemplateApiServerMain:    "",
		TemplateApiServerWire:    "",
	}
}
