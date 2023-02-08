package logic

import (
	"context"
	"douyin/app/user/model"
	"douyin/common/globalkey"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryMessageLogic {
	return &GetHistoryMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取历史消息
func (l *GetHistoryMessageLogic) GetHistoryMessage(in *pb.GetHistoryMessageReq) (*pb.GetHistoryMessageResp, error) {
	follow, err := l.svcCtx.FollowModel.FindOneByUserIdAndToUserId(l.ctx, in.FromUserId, in.ToUserId)
	if follow == nil {
		return nil, errors.Wrapf(ErrUserIsNotFriendError, "当前用户不是好友 userId:%+v,toUserId:%+v", in.FromUserId, in.ToUserId)
	}
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "userid:%v,err:%v", in.FromUserId, err)
	}
	if follow.IsFriend == globalkey.NotMutualAttention {
		return nil, errors.Wrapf(ErrUserIsNotFriendError, "当前用户不是好友 userId:%+v,toUserId:%+v", in.FromUserId, in.ToUserId)
	}
	var query = l.svcCtx.ChatModel.RowBuilder().Where(squirrel.Or{squirrel.Eq{"from_user_id": in.FromUserId, "to_user_id": in.ToUserId}, squirrel.Eq{"to_user_id": in.FromUserId, "from_user_id": in.ToUserId}})
	// TODO 加缓存
	chats, err := l.svcCtx.ChatModel.FindAll(l.ctx, query, "create_time DESC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "userid:%v,err:%v", in.FromUserId, err)
	}
	var res []*pb.Message

	_ = copier.Copy(&res, chats)

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(chats); j++ {
			if res[i].Id == chats[j].Id {
				res[i].CreateTime = chats[j].CreateTime.Format("2006-01-02 15:04:05")
			}
		}
	}

	return &pb.GetHistoryMessageResp{MessageList: res}, nil
}
