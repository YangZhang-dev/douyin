// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: videoRpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VideoRpcClient is the client API for VideoRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoRpcClient interface {
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

type videoRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoRpcClient(cc grpc.ClientConnInterface) VideoRpcClient {
	return &videoRpcClient{cc}
}

func (c *videoRpcClient) UploadVideo(ctx context.Context, in *UploadVideoReq, opts ...grpc.CallOption) (*UploadVideoResp, error) {
	out := new(UploadVideoResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/UploadVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error) {
	out := new(FeedResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/Feed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetFavoriteInfo(ctx context.Context, in *GetFavoriteInfoReq, opts ...grpc.CallOption) (*GetFavoriteInfoResp, error) {
	out := new(GetFavoriteInfoResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/GetFavoriteInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	out := new(PublishListResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/PublishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) Favorite(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error) {
	out := new(FavoriteResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/Favorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) UnFavorite(ctx context.Context, in *UnFavoriteReq, opts ...grpc.CallOption) (*UnFavoriteResp, error) {
	out := new(UnFavoriteResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/UnFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetVideoListByIds(ctx context.Context, in *GetVideoListByIdsReq, opts ...grpc.CallOption) (*GetVideoListByIdsResp, error) {
	out := new(GetVideoListByIdsResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/GetVideoListByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/FavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) PublishComment(ctx context.Context, in *PublishCommentReq, opts ...grpc.CallOption) (*PublishCommentResp, error) {
	out := new(PublishCommentResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/PublishComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error) {
	out := new(DeleteCommentResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListResp, error) {
	out := new(GetCommentListResp)
	err := c.cc.Invoke(ctx, "/pb.videoRpc/GetCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoRpcServer is the server API for VideoRpc service.
// All implementations must embed UnimplementedVideoRpcServer
// for forward compatibility
type VideoRpcServer interface {
	// 发布视频
	UploadVideo(context.Context, *UploadVideoReq) (*UploadVideoResp, error)
	// 视频流
	Feed(context.Context, *FeedReq) (*FeedResp, error)
	// 获取点赞情况
	GetFavoriteInfo(context.Context, *GetFavoriteInfoReq) (*GetFavoriteInfoResp, error)
	// 获取发布列表
	PublishList(context.Context, *PublishListReq) (*PublishListResp, error)
	// 点赞
	Favorite(context.Context, *FavoriteReq) (*FavoriteResp, error)
	// 取消点赞
	UnFavorite(context.Context, *UnFavoriteReq) (*UnFavoriteResp, error)
	// 批量获取视频
	GetVideoListByIds(context.Context, *GetVideoListByIdsReq) (*GetVideoListByIdsResp, error)
	// 获取点赞列表
	FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	// 发布评论
	PublishComment(context.Context, *PublishCommentReq) (*PublishCommentResp, error)
	// 删除评论
	DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentResp, error)
	// 获取视频评论
	GetCommentList(context.Context, *GetCommentListReq) (*GetCommentListResp, error)
	mustEmbedUnimplementedVideoRpcServer()
}

// UnimplementedVideoRpcServer must be embedded to have forward compatible implementations.
type UnimplementedVideoRpcServer struct {
}

func (UnimplementedVideoRpcServer) UploadVideo(context.Context, *UploadVideoReq) (*UploadVideoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedVideoRpcServer) Feed(context.Context, *FeedReq) (*FeedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedVideoRpcServer) GetFavoriteInfo(context.Context, *GetFavoriteInfoReq) (*GetFavoriteInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteInfo not implemented")
}
func (UnimplementedVideoRpcServer) PublishList(context.Context, *PublishListReq) (*PublishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedVideoRpcServer) Favorite(context.Context, *FavoriteReq) (*FavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Favorite not implemented")
}
func (UnimplementedVideoRpcServer) UnFavorite(context.Context, *UnFavoriteReq) (*UnFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavorite not implemented")
}
func (UnimplementedVideoRpcServer) GetVideoListByIds(context.Context, *GetVideoListByIdsReq) (*GetVideoListByIdsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoListByIds not implemented")
}
func (UnimplementedVideoRpcServer) FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedVideoRpcServer) PublishComment(context.Context, *PublishCommentReq) (*PublishCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishComment not implemented")
}
func (UnimplementedVideoRpcServer) DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedVideoRpcServer) GetCommentList(context.Context, *GetCommentListReq) (*GetCommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedVideoRpcServer) mustEmbedUnimplementedVideoRpcServer() {}

// UnsafeVideoRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoRpcServer will
// result in compilation errors.
type UnsafeVideoRpcServer interface {
	mustEmbedUnimplementedVideoRpcServer()
}

func RegisterVideoRpcServer(s grpc.ServiceRegistrar, srv VideoRpcServer) {
	s.RegisterService(&VideoRpc_ServiceDesc, srv)
}

func _VideoRpc_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/UploadVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).UploadVideo(ctx, req.(*UploadVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/Feed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).Feed(ctx, req.(*FeedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetFavoriteInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetFavoriteInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/GetFavoriteInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetFavoriteInfo(ctx, req.(*GetFavoriteInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/PublishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).PublishList(ctx, req.(*PublishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_Favorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).Favorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/Favorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).Favorite(ctx, req.(*FavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_UnFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).UnFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/UnFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).UnFavorite(ctx, req.(*UnFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetVideoListByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoListByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetVideoListByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/GetVideoListByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetVideoListByIds(ctx, req.(*GetVideoListByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/FavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).FavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_PublishComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).PublishComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/PublishComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).PublishComment(ctx, req.(*PublishCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).DeleteComment(ctx, req.(*DeleteCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoRpc/GetCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetCommentList(ctx, req.(*GetCommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoRpc_ServiceDesc is the grpc.ServiceDesc for VideoRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.videoRpc",
	HandlerType: (*VideoRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadVideo",
			Handler:    _VideoRpc_UploadVideo_Handler,
		},
		{
			MethodName: "Feed",
			Handler:    _VideoRpc_Feed_Handler,
		},
		{
			MethodName: "GetFavoriteInfo",
			Handler:    _VideoRpc_GetFavoriteInfo_Handler,
		},
		{
			MethodName: "PublishList",
			Handler:    _VideoRpc_PublishList_Handler,
		},
		{
			MethodName: "Favorite",
			Handler:    _VideoRpc_Favorite_Handler,
		},
		{
			MethodName: "UnFavorite",
			Handler:    _VideoRpc_UnFavorite_Handler,
		},
		{
			MethodName: "GetVideoListByIds",
			Handler:    _VideoRpc_GetVideoListByIds_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _VideoRpc_FavoriteList_Handler,
		},
		{
			MethodName: "PublishComment",
			Handler:    _VideoRpc_PublishComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _VideoRpc_DeleteComment_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _VideoRpc_GetCommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "videoRpc.proto",
}