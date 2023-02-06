package logic

import (
	"context"
	"douyin/app/video/model"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FavoriteList 获取点赞列表
func (l *FavoriteListLogic) FavoriteList(in *pb.FavoriteListReq) (*pb.FavoriteListResp, error) {
	var query = l.svcCtx.FavoriteModel.RowBuilder().Where(squirrel.Eq{"user_id": in.UserId})
	// TODO 加缓存
	favoriteVideos, err := l.svcCtx.FavoriteModel.FindAll(l.ctx, query, "create_time  DESC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "user_id:%v,err:%v", in.UserId, err)
	}
	var ids []int64
	for _, video := range favoriteVideos {
		ids = append(ids, video.VideoId)
	}
	getVideoListByIdsLogic := NewGetVideoListByIdsLogic(l.ctx, l.svcCtx)
	videos, err := getVideoListByIdsLogic.GetVideoListByIds(&pb.GetVideoListByIdsReq{
		UserId: in.CurUserId,
		Ids:    ids,
	})
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Video, len(videos.VideoList))
	_ = copier.Copy(&res, videos.VideoList)

	return &pb.FavoriteListResp{VideoList: res}, nil
}
