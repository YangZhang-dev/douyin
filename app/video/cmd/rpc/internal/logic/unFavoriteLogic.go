package logic

import (
	"context"
	"douyin/app/video/model"
	"douyin/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrVideoNotExistError = xerr.NewErrCode(xerr.VIDEO_NOT_EXIST_ERROR)
var ErrVideoNotFavoriteError = xerr.NewErrCode(xerr.VIDEO_NOT_FAVORITE_ERROR)

type UnFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFavoriteLogic {
	return &UnFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UnFavorite 取消点赞
func (l *UnFavoriteLogic) UnFavorite(in *pb.UnFavoriteReq) (*pb.UnFavoriteResp, error) {
	userId := in.UserId
	videoId := in.VideoId

	video, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "videoId:%v,err:%v", videoId, err)
	}
	if video == nil {
		return nil, errors.Wrapf(ErrVideoNotExistError, "视频不存在 video_id:%+v", videoId)
	}
	favoriteVideo, err := l.svcCtx.FavoriteModel.FindOneByUserIdAndVideoId(l.ctx, userId, videoId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "videoId:%v,userid:%+v,err:%v", videoId, userId, err)
	}

	err = l.svcCtx.VideoModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {

		if favoriteVideo == nil {
			return errors.Wrapf(ErrVideoNotFavoriteError, "视频未点赞 user_id:%+v,video_id:%+v", userId, videoId)
		}
		err := l.svcCtx.FavoriteModel.DeleteByUserIdAndVideoId(l.ctx, session, userId, videoId)
		if err != nil {
			return errors.Wrapf(ErrDBError, "user_id:%+v videoId:%v,err:%v", userId, videoId, err)
		}
		video.FavoriteCount--
		_, err = l.svcCtx.VideoModel.Update(l.ctx, session, video)
		if err != nil {
			return errors.Wrapf(ErrDBError, "videoId:%v,err:%v", videoId, err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.UnFavoriteResp{}, nil
}
