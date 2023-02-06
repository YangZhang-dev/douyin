package main

import (
	"douyin/app/file/cmd/rpc/internal/config"
	"douyin/app/file/cmd/rpc/internal/server"
	"douyin/app/file/cmd/rpc/internal/svc"
	"douyin/app/file/cmd/rpc/pb"
	"douyin/common/interceptor/rpcserver"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/fileRpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterFileRpcServer(grpcServer, server.NewFileRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	s.AddOptions(grpc.MaxRecvMsgSize(32 << 20))
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
