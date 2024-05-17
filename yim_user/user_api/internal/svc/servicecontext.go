package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"yim_server/core"
	"yim_server/yim_user/user_api/internal/config"
	"yim_server/yim_user/user_rpc/types/user_rpc"
	"yim_server/yim_user/user_rpc/users"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user_rpc.UsersClient
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := core.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		DB:      mysqlDB,
	}
}
