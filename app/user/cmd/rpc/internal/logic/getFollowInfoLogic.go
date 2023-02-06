package logic

import (
	"context"
	"douyin/app/user/model"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowInfoLogic {
	return &GetFollowInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// follow
func (l *GetFollowInfoLogic) GetFollowInfo(in *pb.GetFollowInfoReq) (*pb.GetFollowInfoResp, error) {

	follow, err := l.svcCtx.FollowModel.FindOneByUserIdAndToUserId(l.ctx, in.UserId, in.ToUserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "user_id:%v,to_user_id:%v,err:%v", in.UserId, in.ToUserId, err)
	}
	if follow == nil {
		return &pb.GetFollowInfoResp{IsFollow: false}, nil
	}
	var res pb.Follow
	_ = copier.Copy(&res, follow)
	return &pb.GetFollowInfoResp{IsFollow: true, Follow: &res}, nil
}
