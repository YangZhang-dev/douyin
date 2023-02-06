package logic

import (
	"context"
	"douyin/app/video/model"
	"github.com/pkg/errors"

	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteInfoLogic {
	return &GetFavoriteInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFavoriteInfo 获取点赞情况
func (l *GetFavoriteInfoLogic) GetFavoriteInfo(in *pb.GetFavoriteInfoReq) (*pb.GetFavoriteInfoResp, error) {

	favorite, err := l.svcCtx.FavoriteModel.FindOneByUserIdAndVideoId(l.ctx, in.UserId, in.VideoId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "user_id:%v,video_id:%v,err:%v", in.UserId, in.VideoId, err)
	}
	if favorite == nil {
		return &pb.GetFavoriteInfoResp{IsFavorite: false}, nil
	}
	return &pb.GetFavoriteInfoResp{IsFavorite: true}, nil
}
