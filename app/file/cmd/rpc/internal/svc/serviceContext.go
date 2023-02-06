package svc

import (
	"douyin/app/file/cmd/rpc/internal/config"
	"douyin/app/file/cmd/rpc/internal/xoss"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ServiceContext struct {
	Config config.Config
	Xoss   *oss.Bucket
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Xoss:   xoss.NewOssClient(c.Xoss),
	}
}
