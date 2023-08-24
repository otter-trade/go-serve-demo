package svc

import (
	"github.com/otter-trade/go-serve-demo/exchange-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.Log)

	// initialize redis
	rds := c.RedisConf.NewRedis()

	return &ServiceContext{
		Config: c,
		Redis:  rds,
	}
}
