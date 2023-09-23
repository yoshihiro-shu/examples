package svc

import (
	"github.com/yoshihiro-shu/examples/grpc/go-zero/greet/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
