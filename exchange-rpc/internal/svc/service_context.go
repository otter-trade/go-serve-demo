package svc

import (
	"github.com/otter-trade/go-serve-demo/exchange-rpc/internal/config"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis

	AdminModel model.AdminModel
	UserModel  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.Log)

	// initialize redis
	rds := c.RedisConf.NewRedis()

	return &ServiceContext{
		Config:     c,
		Redis:      rds,
		AdminModel: model.NewAdminModel(c.MongoUri.Uri, c.MongoUri.Db, "admin"),
		UserModel:  model.NewUserModel(c.MongoUri.Uri, c.MongoUri.Db, "user"),
	}
}
