package logic

import (
	"context"
	"github.com/otter-trade/go-serve-demo/common/i18n"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/internal/svc"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/model"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/pb"
	"google.golang.org/grpc/status"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminAddLogic {
	return &AdminAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------管理员-----------------------
func (l *AdminAddLogic) AdminAdd(in *pb.AdminAddReq) (resp *pb.AdminAddResp, err error) {
	l.Infow("AdminAdd start", logx.Field("in", in))
	resp = &pb.AdminAddResp{}
	// 插入文档
	admin := &model.Admin{
		Name:     "John",
		Email:    "kim@qq.com",
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}
	err = l.svcCtx.AdminModel.Insert(l.ctx, admin)
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		err = status.Error(i18n.DbError, err.Error())
		return
	}

	// 查询文档
	result, err := l.svcCtx.AdminModel.FindOne(l.ctx, "64e73048a9d5231449c34f89")
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		err = status.Error(i18n.DbError, err.Error())
		return
	}

	l.Infow("result", logx.Field("result", result))

	// 更新文档
	update := &model.Admin{
		ID:    result.ID,
		Name:  result.Name,
		Email: "888@qq.com",
	}
	_, err = l.svcCtx.AdminModel.Update(l.ctx, update)
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		err = status.Error(i18n.DbError, err.Error())
		return
	}

	// 删除文档
	_, err = l.svcCtx.AdminModel.Delete(l.ctx, "64e73048a9d5231449c34f89")
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		err = status.Error(i18n.DbError, err.Error())
		return
	}

	return resp, nil
}
