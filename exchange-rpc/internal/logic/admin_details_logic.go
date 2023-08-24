package logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

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

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(l.svcCtx.Config.MongoUri.Uri).SetMaxPoolSize(200))
	if err != nil {
		l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
		return
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			l.Errorw("AdminAdd", logx.Field("in", in), logx.Field("err", err))
			return
		}
	}()

	// 选择数据库和集合
	collection := client.Database(l.svcCtx.Config.MongoUri.Db).Collection("admin")

	objectID, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, err
	}

	result := collection.FindOne(
		context.TODO(), bson.M{"_id": objectID})
	if result.Err() != nil {
		return nil, result.Err()
	}

	u := new(models.Admin)
	if err := result.Decode(u); err != nil {
		return nil, err
	}
	return &pb.AdminDetailsResp{
		Id:        in.Id,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreateAt.Format(time.DateTime),
		UpdatedAt: u.CreateAt.Format(time.DateTime),
	}, nil

	return
}
