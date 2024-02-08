package svc

import (
	gofkConf "github.com/bhmy-shm/gofks/core/config"
)

type ServiceContext struct {
	Config *gofkConf.Config
}

func NewServiceContext(c *gofkConf.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
