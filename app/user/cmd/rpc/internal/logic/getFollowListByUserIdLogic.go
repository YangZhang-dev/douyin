package logic

import (
	"context"
	"douyin/app/user/model"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListByUserIdLogic {
	return &GetFollowListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListByUserIdLogic) GetFollowListByUserId(in *pb.GetFollowListByUserIdReq) (*pb.GetFollowListByUserIdResp, error) {
	// 1. 检查用户是否存在
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "userid:%v,err:%v", in.UserId, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNotExistError, "用户不存在 userid:%v,err:%v", in.UserId, err)
	}
	// 2. 获取关注用户id
	var query squirrel.SelectBuilder
	query = l.svcCtx.FollowModel.RowBuilder().Where(squirrel.Eq{"user_id": in.UserId})
	// TODO 加缓存
	followUsers, err := l.svcCtx.FollowModel.FindAll(l.ctx, query, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "查找关注用户 id:%+v,err:%v", in.UserId, err)
	}

	// 2. 封装用户信息和关注情况
	var ids []int64
	for _, follow := range followUsers {
		ids = append(ids, follow.ToUserId)
	}
	getUserListByIdsLogic := NewGetUserListByIdsLogic(l.ctx, l.svcCtx)
	getUserListByIdsResp, err := getUserListByIdsLogic.GetUserListByIds(&pb.GetUserListByIdsReq{Ids: ids, UserId: in.CurUserId})
	if err != nil {
		return nil, err
	}

	return &pb.GetFollowListByUserIdResp{Users: getUserListByIdsResp.Users}, nil
}
