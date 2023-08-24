package jobs

import (
	"context"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/svc"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type JobService struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobService(ctx context.Context, svcCtx *svc.ServiceContext) *JobService {
	return &JobService{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
	}
}

func (l *JobService) Register() {
	go l.SendMessage()
}

func (l *JobService) SendMessage() {
	crontab := cron.New(cron.WithSeconds()) //精确到秒

	crontab.AddFunc("0 */10 * * * ?", func() {
		//
		l.Infow(" start", logx.Field("in", time.Now().Format(time.DateTime)))
	})

	// 启动定时器
	crontab.Start()
}
