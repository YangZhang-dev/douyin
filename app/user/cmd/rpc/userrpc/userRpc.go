// Code generated by goctl. DO NOT EDIT.
// Source: userRpc.proto

package userrpc

import (
	"context"

	"douyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Follow                    = pb.Follow
	FollowReq                 = pb.FollowReq
	FollowResp                = pb.FollowResp
	GenerateTokenReq          = pb.GenerateTokenReq
	GenerateTokenResp         = pb.GenerateTokenResp
	GetFansListByUserIdReq    = pb.GetFansListByUserIdReq
	GetFansListByUserIdResp   = pb.GetFansListByUserIdResp
	GetFollowInfoReq          = pb.GetFollowInfoReq
	GetFollowInfoResp         = pb.GetFollowInfoResp
	GetFollowListByUserIdReq  = pb.GetFollowListByUserIdReq
	GetFollowListByUserIdResp = pb.GetFollowListByUserIdResp
	GetFriendsListReq         = pb.GetFriendsListReq
	GetFriendsListResp        = pb.GetFriendsListResp
	GetHistoryMessageReq      = pb.GetHistoryMessageReq
	GetHistoryMessageResp     = pb.GetHistoryMessageResp
	GetUserListByIdsReq       = pb.GetUserListByIdsReq
	GetUserListByIdsResp      = pb.GetUserListByIdsResp
	LoginReq                  = pb.LoginReq
	LoginResp                 = pb.LoginResp
	Message                   = pb.Message
	RegisterReq               = pb.RegisterReq
	RegisterResp              = pb.RegisterResp
	SendMessageReq            = pb.SendMessageReq
	SendMessageResp           = pb.SendMessageResp
	UnFollowReq               = pb.UnFollowReq
	UnFollowResp              = pb.UnFollowResp
	User                      = pb.User

	UserRpc interface {
		// --------------------------------user--------------------------------
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		// 用户登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 批量获取用户信息
		GetUserListByIds(ctx context.Context, in *GetUserListByIdsReq, opts ...grpc.CallOption) (*GetUserListByIdsResp, error)
		// 获取Token
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		// --------------------------------follow--------------------------------
		GetFollowInfo(ctx context.Context, in *GetFollowInfoReq, opts ...grpc.CallOption) (*GetFollowInfoResp, error)
		// 获取关注列表
		GetFollowListByUserId(ctx context.Context, in *GetFollowListByUserIdReq, opts ...grpc.CallOption) (*GetFollowListByUserIdResp, error)
		// 获取粉丝列表
		GetFansListByUserId(ctx context.Context, in *GetFansListByUserIdReq, opts ...grpc.CallOption) (*GetFansListByUserIdResp, error)
		// 关注
		Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error)
		// 取消关注
		UnFollow(ctx context.Context, in *UnFollowReq, opts ...grpc.CallOption) (*UnFollowResp, error)
		// 获取好友列表
		GetFriendsList(ctx context.Context, in *GetFriendsListReq, opts ...grpc.CallOption) (*GetFriendsListResp, error)
		// --------------------------------chat--------------------------------
		SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
		// 获取历史消息
		GetHistoryMessage(ctx context.Context, in *GetHistoryMessageReq, opts ...grpc.CallOption) (*GetHistoryMessageResp, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

// --------------------------------user--------------------------------
func (m *defaultUserRpc) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 用户登录
func (m *defaultUserRpc) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 批量获取用户信息
func (m *defaultUserRpc) GetUserListByIds(ctx context.Context, in *GetUserListByIdsReq, opts ...grpc.CallOption) (*GetUserListByIdsResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetUserListByIds(ctx, in, opts...)
}

// 获取Token
func (m *defaultUserRpc) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

// --------------------------------follow--------------------------------
func (m *defaultUserRpc) GetFollowInfo(ctx context.Context, in *GetFollowInfoReq, opts ...grpc.CallOption) (*GetFollowInfoResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetFollowInfo(ctx, in, opts...)
}

// 获取关注列表
func (m *defaultUserRpc) GetFollowListByUserId(ctx context.Context, in *GetFollowListByUserIdReq, opts ...grpc.CallOption) (*GetFollowListByUserIdResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetFollowListByUserId(ctx, in, opts...)
}

// 获取粉丝列表
func (m *defaultUserRpc) GetFansListByUserId(ctx context.Context, in *GetFansListByUserIdReq, opts ...grpc.CallOption) (*GetFansListByUserIdResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetFansListByUserId(ctx, in, opts...)
}

// 关注
func (m *defaultUserRpc) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

// 取消关注
func (m *defaultUserRpc) UnFollow(ctx context.Context, in *UnFollowReq, opts ...grpc.CallOption) (*UnFollowResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UnFollow(ctx, in, opts...)
}

// 获取好友列表
func (m *defaultUserRpc) GetFriendsList(ctx context.Context, in *GetFriendsListReq, opts ...grpc.CallOption) (*GetFriendsListResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetFriendsList(ctx, in, opts...)
}

// --------------------------------chat--------------------------------
func (m *defaultUserRpc) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SendMessage(ctx, in, opts...)
}

// 获取历史消息
func (m *defaultUserRpc) GetHistoryMessage(ctx context.Context, in *GetHistoryMessageReq, opts ...grpc.CallOption) (*GetHistoryMessageResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetHistoryMessage(ctx, in, opts...)
}