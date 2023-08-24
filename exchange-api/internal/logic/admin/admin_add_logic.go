package admin

import (
	"context"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/exchange"
	"net/http"

	"github.com/otter-trade/go-serve-demo/exchange-api/internal/svc"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminAddLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminAddLogic(r *http.Request, svcCtx *svc.ServiceContext) *AdminAddLogic {
	return &AdminAddLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *AdminAddLogic) AdminAdd(req *types.AdminAddReq) (resp *types.AdminCommonResp, err error) {
	l.Infow("AdminAdd start", logx.Field("req", req))
	resp = &types.AdminCommonResp{}

	_, err = l.svcCtx.Exchange.AdminAdd(l.ctx, &exchange.AdminAddReq{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	return
}
