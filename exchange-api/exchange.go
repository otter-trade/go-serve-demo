package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/jobs"

	"github.com/otter-trade/go-serve-demo/exchange-api/internal/config"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/handler"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/exchange.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//注册JOB
	scheduler := jobs.NewJobService(context.Background(), svcCtx)
	scheduler.Register()

	handler.RegisterHandlers(server, svcCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
