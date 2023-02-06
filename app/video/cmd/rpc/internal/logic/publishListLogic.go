package logic

import (
	"context"
	"douyin/app/user/model"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PublishList 获取发布列表
func (l *PublishListLogic) PublishList(in *pb.PublishListReq) (*pb.PublishListResp, error) {

	var query = l.svcCtx.VideoModel.RowBuilder().Where(squirrel.Eq{"user_id": in.UserId})
	// TODO 加缓存
	videos, err := l.svcCtx.VideoModel.FindAll(l.ctx, query, "create_time  DESC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "user_id:%v,err:%v", in.UserId, err)
	}

	res := make([]*pb.Video, len(videos))
	_ = copier.Copy(&res, videos)

	return &pb.PublishListResp{VideoList: res}, nil
}
