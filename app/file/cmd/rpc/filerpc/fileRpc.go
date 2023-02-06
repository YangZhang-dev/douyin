// Code generated by goctl. DO NOT EDIT.
// Source: fileRpc.proto

package filerpc

import (
	"context"

	"douyin/app/file/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UploadVideoByLocalReq  = pb.UploadVideoByLocalReq
	UploadVideoByLocalResp = pb.UploadVideoByLocalResp
	UploadVideoByOssReq    = pb.UploadVideoByOssReq
	UploadVideoByOssResp   = pb.UploadVideoByOssResp

	FileRpc interface {
		// 上传视频到OSS
		UploadVideoByOss(ctx context.Context, in *UploadVideoByOssReq, opts ...grpc.CallOption) (*UploadVideoByOssResp, error)
		// 上传视频到本地
		UploadVideoByLocal(ctx context.Context, in *UploadVideoByLocalReq, opts ...grpc.CallOption) (*UploadVideoByLocalResp, error)
	}

	defaultFileRpc struct {
		cli zrpc.Client
	}
)

func NewFileRpc(cli zrpc.Client) FileRpc {
	return &defaultFileRpc{
		cli: cli,
	}
}

// 上传视频到OSS
func (m *defaultFileRpc) UploadVideoByOss(ctx context.Context, in *UploadVideoByOssReq, opts ...grpc.CallOption) (*UploadVideoByOssResp, error) {
	client := pb.NewFileRpcClient(m.cli.Conn())
	return client.UploadVideoByOss(ctx, in, opts...)
}

// 上传视频到本地
func (m *defaultFileRpc) UploadVideoByLocal(ctx context.Context, in *UploadVideoByLocalReq, opts ...grpc.CallOption) (*UploadVideoByLocalResp, error) {
	client := pb.NewFileRpcClient(m.cli.Conn())
	return client.UploadVideoByLocal(ctx, in, opts...)
}
