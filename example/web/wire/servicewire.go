package wire

import (
	gofkConf "github.com/bhmy-shm/gofks/core/config"
	"web/svc"
)

type ServiceWire struct {
	Ctx *svc.ServiceContext
}

func NewServiceWire(c *gofkConf.Config) *ServiceWire {
	return &ServiceWire{
		Ctx: svc.NewServiceContext(c),
	}
}

func (s *ServiceWire) ServiceCtx() *svc.ServiceContext {
	return s.Ctx
}
