package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"qql/apps/im/api/internal/config"
	"qql/apps/im/rpc/im"
	"qql/apps/im/rpc/imclient"
	"qql/apps/social/rpc/social"
	"qql/apps/social/rpc/socialclient"
	"qql/apps/user/rpc/user"
	"qql/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	im.ImClient
	social.SocialClient
	user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ImClient:     imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
		SocialClient: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		UserClient:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
