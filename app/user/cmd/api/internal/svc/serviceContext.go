package svc

import (
	"douyin/app/user/cmd/api/internal/config"
	"douyin/app/user/cmd/rpc/userrpc"
	"douyin/common/middleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	OptionalJWTMiddleware rest.Middleware
	JwtAuthMiddleWare     rest.Middleware
	UserRpc               userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpc := userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf))

	return &ServiceContext{
		Config:                c,
		OptionalJWTMiddleware: middleware.NewOptionalJWTMiddleware(c.JwtAuth.AccessSecret).Handle,
		JwtAuthMiddleWare:     middleware.NewJwtAuthMiddleWareMiddleware(c.JwtAuth.AccessSecret).Handle,
		UserRpc:               userRpc,
	}
}
