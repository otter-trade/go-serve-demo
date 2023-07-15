// Code generated by goctl. DO NOT EDIT.
// Source: exchange.proto

package exchange

import (
	"context"

	"github.com/otter-trade/go-serve-demo/exchange-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Access            = pb.Access
	AccessAddReq      = pb.AccessAddReq
	AccessAddResp     = pb.AccessAddResp
	AccessDeleteReq   = pb.AccessDeleteReq
	AccessDeleteResp  = pb.AccessDeleteResp
	AccessDetailsReq  = pb.AccessDetailsReq
	AccessDetailsResp = pb.AccessDetailsResp
	AccessListReq     = pb.AccessListReq
	AccessListResp    = pb.AccessListResp
	AccessSearchReq   = pb.AccessSearchReq
	AccessSearchResp  = pb.AccessSearchResp
	AccessUpdateReq   = pb.AccessUpdateReq
	AccessUpdateResp  = pb.AccessUpdateResp
	Admin             = pb.Admin
	AdminAddReq       = pb.AdminAddReq
	AdminAddResp      = pb.AdminAddResp
	AdminDeleteReq    = pb.AdminDeleteReq
	AdminDeleteResp   = pb.AdminDeleteResp
	AdminDetailsReq   = pb.AdminDetailsReq
	AdminDetailsResp  = pb.AdminDetailsResp
	AdminSearchReq    = pb.AdminSearchReq
	AdminSearchResp   = pb.AdminSearchResp
	AdminUpdateReq    = pb.AdminUpdateReq
	AdminUpdateResp   = pb.AdminUpdateResp

	Exchange interface {
		// -----------------------管理员-----------------------
		AdminAdd(ctx context.Context, in *AdminAddReq, opts ...grpc.CallOption) (*AdminAddResp, error)
		// 更新管理员
		AdminUpdate(ctx context.Context, in *AdminUpdateReq, opts ...grpc.CallOption) (*AdminUpdateResp, error)
		// 删除管理员
		AdminDelete(ctx context.Context, in *AdminDeleteReq, opts ...grpc.CallOption) (*AdminDeleteResp, error)
		// 获取管理员
		AdminDetails(ctx context.Context, in *AdminDetailsReq, opts ...grpc.CallOption) (*AdminDetailsResp, error)
		// 管理员列表
		AdminSearch(ctx context.Context, in *AdminSearchReq, opts ...grpc.CallOption) (*AdminSearchResp, error)
	}

	defaultExchange struct {
		cli zrpc.Client
	}
)

func NewExchange(cli zrpc.Client) Exchange {
	return &defaultExchange{
		cli: cli,
	}
}

// -----------------------管理员-----------------------
func (m *defaultExchange) AdminAdd(ctx context.Context, in *AdminAddReq, opts ...grpc.CallOption) (*AdminAddResp, error) {
	client := pb.NewExchangeClient(m.cli.Conn())
	return client.AdminAdd(ctx, in, opts...)
}

// 更新管理员
func (m *defaultExchange) AdminUpdate(ctx context.Context, in *AdminUpdateReq, opts ...grpc.CallOption) (*AdminUpdateResp, error) {
	client := pb.NewExchangeClient(m.cli.Conn())
	return client.AdminUpdate(ctx, in, opts...)
}

// 删除管理员
func (m *defaultExchange) AdminDelete(ctx context.Context, in *AdminDeleteReq, opts ...grpc.CallOption) (*AdminDeleteResp, error) {
	client := pb.NewExchangeClient(m.cli.Conn())
	return client.AdminDelete(ctx, in, opts...)
}

// 获取管理员
func (m *defaultExchange) AdminDetails(ctx context.Context, in *AdminDetailsReq, opts ...grpc.CallOption) (*AdminDetailsResp, error) {
	client := pb.NewExchangeClient(m.cli.Conn())
	return client.AdminDetails(ctx, in, opts...)
}

// 管理员列表
func (m *defaultExchange) AdminSearch(ctx context.Context, in *AdminSearchReq, opts ...grpc.CallOption) (*AdminSearchResp, error) {
	client := pb.NewExchangeClient(m.cli.Conn())
	return client.AdminSearch(ctx, in, opts...)
}
