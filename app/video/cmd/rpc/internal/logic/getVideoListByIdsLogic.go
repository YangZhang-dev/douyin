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

type GetVideoListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListByIdsLogic {
	return &GetVideoListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetVideoListByIds 批量获取视频
func (l *GetVideoListByIdsLogic) GetVideoListByIds(in *pb.GetVideoListByIdsReq) (*pb.GetVideoListByIdsResp, error) {
	ids := in.Ids
	var query squirrel.SelectBuilder

	query = l.svcCtx.VideoModel.RowBuilder().Where(squirrel.Eq{"id": ids})
	// TODO 加缓存
	videos, err := l.svcCtx.VideoModel.FindAll(l.ctx, query, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "通过ids查找视频 ids:%+v,err:%v", ids, err)
	}
	//res := make([]*pb.Video, len(videos))
	var res []*pb.Video
	_ = copier.Copy(&res, videos)

	if in.UserId != nil {
		id := *in.UserId
		getFollowInfoLogic := NewGetFavoriteInfoLogic(l.ctx, l.svcCtx)
		for _, video := range res {
			favoriteInfoResp, err := getFollowInfoLogic.GetFavoriteInfo(&pb.GetFavoriteInfoReq{
				UserId:  id,
				VideoId: video.Id,
			})
			if err != nil {
				return nil, err
			}
			video.IsFavorite = favoriteInfoResp.IsFavorite
		}
	}
	return &pb.GetVideoListByIdsResp{VideoList: res}, nil
}
