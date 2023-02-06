package logic

import (
	"context"
	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/app/user/model"
	"douyin/common/globalkey"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserIsNotFriendError = xerr.NewErrCode(xerr.USER_IS_NOT_FRIEND_ERROR)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *pb.SendMessageReq) (*pb.SendMessageResp, error) {

	follow, err := l.svcCtx.FollowModel.FindOneByUserIdAndToUserId(l.ctx, in.UserId, in.ToUserId)
	if follow == nil {
		return nil, errors.Wrapf(ErrUserIsNotFriendError, "当前用户不是好友 userId:%+v,toUserId:%+v", in.UserId, in.ToUserId)
	}
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "userid:%v,err:%v", in.UserId, err)
	}
	if follow.IsFriend == globalkey.NotMutualAttention {
		return nil, errors.Wrapf(ErrUserIsNotFriendError, "当前用户不是好友 userId:%+v,toUserId:%+v", in.UserId, in.ToUserId)
	}
	_, err = l.svcCtx.ChatModel.Insert(l.ctx, nil, &model.Chat{
		UserId:   in.UserId,
		ToUserId: in.ToUserId,
		Content:  in.Content,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrDBError, "userid:%v,err:%v", in.UserId, err)
	}

	return &pb.SendMessageResp{}, nil
}
