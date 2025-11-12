package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"qql/apps/user/api/internal/config"
	"qql/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	*redis.Redis
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Redis: redis.MustNewRedis(c.Redisx),
		User:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
