package svc

import "qql/apps/task/internal/config"

type ServiceContext struct {
	config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{c}
}
