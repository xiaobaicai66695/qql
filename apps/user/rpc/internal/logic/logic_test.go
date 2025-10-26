package logic

import (
	"github.com/zeromicro/go-zero/core/conf"
	"path/filepath"
	"qql/apps/user/rpc/internal/config"
	"qql/apps/user/rpc/internal/svc"
)

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad(filepath.Join("../../etc/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}
