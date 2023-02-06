// Code generated by goctl. DO NOT EDIT.
// Source: videoRpc.proto

package videorpc

import (
	"context"

	"douyin/app/video/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment               = pb.Comment
	DeleteCommentReq      = pb.DeleteCommentReq
	DeleteCommentResp     = pb.DeleteCommentResp
	FavoriteListReq       = pb.FavoriteListReq
	FavoriteListResp      = pb.FavoriteListResp
	FavoriteReq           = pb.FavoriteReq
	FavoriteResp          = pb.FavoriteResp
	FeedReq               = pb.FeedReq
	FeedResp              = pb.FeedResp
	GetCommentListReq     = pb.GetCommentListReq
	GetCommentListResp    = pb.GetCommentListResp
	GetFavoriteInfoReq    = pb.GetFavoriteInfoReq
	GetFavoriteInfoResp   = pb.GetFavoriteInfoResp
	GetVideoListByIdsReq  = pb.GetVideoListByIdsReq
	GetVideoListByIdsResp = pb.GetVideoListByIdsResp
	PublishCommentReq     = pb.PublishCommentReq
	PublishCommentResp    = pb.PublishCommentResp
	PublishListReq        = pb.PublishListReq
	PublishListResp       = pb.PublishListResp
	UnFavoriteReq         = pb.UnFavoriteReq
	UnFavoriteResp        = pb.UnFavoriteResp
	UploadVideoReq        = pb.UploadVideoReq
	UploadVideoResp       = pb.UploadVideoResp
	Video                 = pb.Video

	VideoRpc interface {
		// 发布视频
		UploadVideo(ctx context.Context, in *UploadVideoReq, opts ...grpc.CallOption) (*UploadVideoResp, error)
		// 视频流
		Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error)
		// 获取点赞情况
		GetFavoriteInfo(ctx context.Context, in *GetFavoriteInfoReq, opts ...grpc.CallOption) (*GetFavoriteInfoResp, error)
		// 获取发布列表
		PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
		// 点赞
		Favorite(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error)
		// 取消点赞
		UnFavorite(ctx context.Context, in *UnFavoriteReq, opts ...grpc.CallOption) (*UnFavoriteResp, error)
		// 批量获取视频
		GetVideoListByIds(ctx context.Context, in *GetVideoListByIdsReq, opts ...grpc.CallOption) (*GetVideoListByIdsResp, error)
		// 获取点赞列表
		FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
		// 发布评论
		PublishComment(ctx context.Context, in *PublishCommentReq, opts ...grpc.CallOption) (*PublishCommentResp, error)
		// 删除评论
		DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error)
		// 获取视频评论
		GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListResp, error)
	}

	defaultVideoRpc struct {
		cli zrpc.Client
	}
)

func NewVideoRpc(cli zrpc.Client) VideoRpc {
	return &defaultVideoRpc{
		cli: cli,
	}
}

// 发布视频
func (m *defaultVideoRpc) UploadVideo(ctx context.Context, in *UploadVideoReq, opts ...grpc.CallOption) (*UploadVideoResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.UploadVideo(ctx, in, opts...)
}

// 视频流
func (m *defaultVideoRpc) Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.Feed(ctx, in, opts...)
}

// 获取点赞情况
func (m *defaultVideoRpc) GetFavoriteInfo(ctx context.Context, in *GetFavoriteInfoReq, opts ...grpc.CallOption) (*GetFavoriteInfoResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.GetFavoriteInfo(ctx, in, opts...)
}

// 获取发布列表
func (m *defaultVideoRpc) PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.PublishList(ctx, in, opts...)
}

// 点赞
func (m *defaultVideoRpc) Favorite(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.Favorite(ctx, in, opts...)
}

// 取消点赞
func (m *defaultVideoRpc) UnFavorite(ctx context.Context, in *UnFavoriteReq, opts ...grpc.CallOption) (*UnFavoriteResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.UnFavorite(ctx, in, opts...)
}

// 批量获取视频
func (m *defaultVideoRpc) GetVideoListByIds(ctx context.Context, in *GetVideoListByIdsReq, opts ...grpc.CallOption) (*GetVideoListByIdsResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.GetVideoListByIds(ctx, in, opts...)
}

// 获取点赞列表
func (m *defaultVideoRpc) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.FavoriteList(ctx, in, opts...)
}

// 发布评论
func (m *defaultVideoRpc) PublishComment(ctx context.Context, in *PublishCommentReq, opts ...grpc.CallOption) (*PublishCommentResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.PublishComment(ctx, in, opts...)
}

// 删除评论
func (m *defaultVideoRpc) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

// 获取视频评论
func (m *defaultVideoRpc) GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListResp, error) {
	client := pb.NewVideoRpcClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}
