package admin

import (
	"context"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/exchange"
	"net/http"

	"github.com/otter-trade/go-serve-demo/exchange-api/internal/svc"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDetailLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDetailLogic(r *http.Request, svcCtx *svc.ServiceContext) *AdminDetailLogic {
	return &AdminDetailLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *AdminDetailLogic) AdminDetail(req *types.AdminDetailReq) (resp *types.AdminDetailResp, err error) {
	l.Infow("AdminDetail start", logx.Field("req", req))

	resp = &types.AdminDetailResp{}
	_, err = l.svcCtx.Exchange.AdminDetails(l.ctx, &exchange.AdminDetailsReq{Id: req.Id})
	if err != nil {
		l.Errorw("AdminDetail", logx.Field("req", req), logx.Field("err", err))
		return
	}

	return
}
