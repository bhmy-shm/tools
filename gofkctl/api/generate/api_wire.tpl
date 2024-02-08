package wire

import (
	{{.svcImport}}
    {{.configImport}}
)

type ServiceWire struct {
	Ctx  *svc.ServiceContext
}

func NewServiceWire(c *{{.config}}) *ServiceWire {
	return &ServiceWire{
	    Ctx: svc.NewServiceContext(c),
	}
}

func (s *ServiceWire) ServiceCtx() *svc.ServiceContext {
	return s.Ctx
}
