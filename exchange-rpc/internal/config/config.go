package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	MongoUri  MongoUri
}

type MongoUri struct {
	Uri      string
	Db       string
	PoolSize uint64
}
