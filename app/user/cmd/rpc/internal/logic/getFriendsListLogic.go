package logic

import (
	"context"
	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsListLogic {
	return &GetFriendsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取好友列表
func (l *GetFriendsListLogic) GetFriendsList(in *pb.GetFriendsListReq) (*pb.GetFriendsListResp, error) {

	getFollowListByUserIdLogic := NewGetFollowListByUserIdLogic(l.ctx, l.svcCtx)
	getFansListByUserIdLogic := NewGetFansListByUserIdLogic(l.ctx, l.svcCtx)

	followListByUserIdResp, err := getFollowListByUserIdLogic.GetFollowListByUserId(&pb.GetFollowListByUserIdReq{
		UserId:    in.UserId,
		CurUserId: in.CurUserId,
	})
	if err != nil {
		return nil, err
	}
	fansListByUserIdResp, err := getFansListByUserIdLogic.GetFansListByUserId(&pb.GetFansListByUserIdReq{
		UserId:    in.UserId,
		CurUserId: in.CurUserId,
	})
	if err != nil {
		return nil, err
	}
	res := getMutualAttention(followListByUserIdResp.Users, fansListByUserIdResp.Users)
	return &pb.GetFriendsListResp{User: res}, nil
}
func getMutualAttention(A, B []*pb.User) []*pb.User {
	if len(A) < 1 || len(B) < 1 {
		return nil
	}
	result := make([]*pb.User, 0)
	// 去重
	flagMap := make(map[int64]bool, 0)
	for _, a := range A {
		flagMap[a.Id] = true
	}
	for _, b := range B {
		if flagMap[b.Id] {
			result = append(result, b)
		}
	}
	return result
}
