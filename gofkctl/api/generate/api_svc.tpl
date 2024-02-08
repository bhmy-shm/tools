package svc

import (
    {{.configImport}}
)

type ServiceContext struct {
	{{.middleware}}
	Config *{{.config}}
}

func NewServiceContext(c *{{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		{{.middlewareAssignment}}
	}
}
