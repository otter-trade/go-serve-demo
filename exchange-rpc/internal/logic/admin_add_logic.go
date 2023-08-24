package logic

import (
	"context"
	"github.com/otter-trade/go-serve-demo/common/helpers"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/internal/models"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/internal/svc"
	"github.com/otter-trade/go-serve-demo/exchange-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
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

	collection := mon.MustNewModel(l.svcCtx.Config.MongoUri.Uri, l.svcCtx.Config.MongoUri.Db, "admin")

	//client, err := mongo.Connect(l.ctx, options.Client().ApplyURI(l.svcCtx.Config.MongoUri.Uri).SetMaxPoolSize(l.svcCtx.Config.MongoUri.PoolSize))
	//if err != nil {
	//	l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
	//	return
	//}
	//defer func() {
	//	if err = client.Disconnect(context.Background()); err != nil {
	//		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
	//		return
	//	}
	//}()

	// 选择数据库和集合
	//collection := client.Database(l.svcCtx.Config.MongoUri.Db).Collection("admin")
	// 插入文档
	admin := models.Admin{
		ID:        helpers.GetUUID(),
		Name:      "John",
		Email:     "kim@qq.com",
		CreatedAt: time.Now().Format(time.DateTime),
		UpdatedAt: time.Now().Format(time.DateTime),
	}
	_, err = collection.InsertOne(l.ctx, admin)
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		return
	}

	// 查询文档
	filter := bson.M{"name": "John"}
	var result models.Admin
	err = collection.FindOne(l.ctx, &result, filter)
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		return
	}

	l.Infow("result", logx.Field("result", result))

	// 更新文档
	update := bson.M{"$set": bson.M{"email": "888@qq.com"}}
	_, err = collection.UpdateOne(l.ctx, bson.M{"_id": result.ID}, update)
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		return
	}

	// 删除文档
	_, err = collection.DeleteOne(l.ctx, filter)
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		return
	}

	return resp, nil
}
