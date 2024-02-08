package pathx

var rpcMethods = `func (this *{{ .StructFieldName }}) {{ .Name }}(req *rpcFlag.{{ .RequestType }}, resp *rpcFlag.{{ .ResponseType }}) error {
	req = new(rpcFlag.{{ .RequestType }})
	resp = new(rpcFlag.{{ .ResponseType }})
	err := this.validator.ValidateCreateRequest(req)
	//err = group.NewGroup().Insert(this.pgsql.DB())
	return err
}`

var services = `package {{ .PackageName }}

import (
	storage "hitry-go/example/agentServer/internal/storage/{{ .ProtoName }}"
	check "hitry-go/example/agentServer/internal/validators/{{ .ProtoName }}"
	rpcFlag "hitry-go/example/agentServer/protoJson/{{ .ProtoName }}"
	"hitry-go/hitry"
)

type  Service{{ .FieldName }} struct {
	validator hitry.Validator
	storage   hitry.Storage
	sql       storage.I{{ .FieldName }}SqlInter
    cache     storage.I{{ .FieldName }}RedisInter
    mq        storage.I{{ .FieldName }}MqInter
}

func New{{ .FieldName }}(sql hitry.Storage) *Service{{ .FieldName }} {

	{{ .FieldName }}Storage := sql.(*storage.Storage{{ .FieldName }})

	return &Service{{ .FieldName }} {
		validator: check.New(),
		storage:   {{ .FieldName }}Storage,
		cache:     {{ .FieldName }}Storage.GetRedis(),
		mq:        {{ .FieldName }}Storage.GetMq(),
	}
}

{{ range .RpcMethods }}
func (this *Service{{ $.FieldName }}) {{ .Name }}(req *rpcFlag.{{ .RequestType }}, resp *rpcFlag.{{ .ResponseType }}) error {
	req = new(rpcFlag.{{ .RequestType }})
	resp = new(rpcFlag.{{ .ResponseType }})
	err := this.validator.ValidateCreateRequest(req)
	//err = group.NewGroup().Insert(this.pgsql.DB())
	return err
}
{{ end }}
`
