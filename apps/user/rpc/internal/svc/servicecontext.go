package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	constants2 "qql/apps/pkg/constants"
	"qql/apps/user/models"
	"qql/apps/user/rpc/internal/config"
	"qql/pkg/ctxdata"
	"time"
)

type ServiceContext struct {
	Config config.Config

	*redis.Redis

	UserModels models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.Redisx),
		UserModels: models.NewUsersModel(sqlConn, c.Cache),
	}
}

func (svc *ServiceContext) SetRootToken() error {
	//生成token
	systemToken, err := ctxdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(), 999999999999, constants2.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}
	return svc.Redis.Set(constants2.REDIS_SYSTEM_ROOT_TOKEN, systemToken)
}
