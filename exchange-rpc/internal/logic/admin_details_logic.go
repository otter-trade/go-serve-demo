package logic

import (
	"context"

	"github.com/otter-trade/go-serve-demo/exchange-rpc/internal/svc"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDetailsLogic {
	return &AdminDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取管理员
func (l *AdminDetailsLogic) AdminDetails(in *pb.AdminDetailsReq) (resp *pb.AdminDetailsResp, err error) {
	l.Infow("AdminDetails start", logx.Field("in", in))
	resp = &pb.AdminDetailsResp{}

	// 查数据、做逻辑

	return
}
