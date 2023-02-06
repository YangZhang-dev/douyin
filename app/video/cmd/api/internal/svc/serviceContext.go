package svc

import (
	"douyin/app/file/cmd/rpc/filerpc"
	"douyin/app/user/cmd/rpc/userrpc"
	"douyin/app/video/cmd/api/internal/config"
	videoMiddleware "douyin/app/video/cmd/api/internal/middleware"
	"douyin/app/video/cmd/rpc/videorpc"
	"douyin/common/middleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config                config.Config
	OptionalJWTMiddleware rest.Middleware
	JwtAuthMiddleWare     rest.Middleware
	ParseFormMiddleware   rest.Middleware
	FileRpc               filerpc.FileRpc
	VideoRpc              videorpc.VideoRpc
	UserRpc               userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 32))
	fileRpc := filerpc.NewFileRpc(zrpc.MustNewClient(c.FileRpcConf, zrpc.WithDialOption(dialOption)))

	videoRpc := videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf))
	userRpc := userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf))
	return &ServiceContext{
		Config:                c,
		OptionalJWTMiddleware: middleware.NewOptionalJWTMiddleware(c.JwtAuth.AccessSecret).Handle,
		JwtAuthMiddleWare:     middleware.NewJwtAuthMiddleWareMiddleware(c.JwtAuth.AccessSecret).Handle,
		ParseFormMiddleware:   videoMiddleware.NewParseFormMiddleware(c.JwtAuth.AccessSecret).Handle,
		FileRpc:               fileRpc,
		VideoRpc:              videoRpc,
		UserRpc:               userRpc,
	}
}
