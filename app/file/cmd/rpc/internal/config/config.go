package config

import (
	"douyin/app/file/cmd/rpc/internal/xoss"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Xoss xoss.Xoss
}
