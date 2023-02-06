package logic

import (
	"context"
	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"
	"douyin/common/globalkey"
	"douyin/common/tool"
	"douyin/common/xerr"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrDataFormatError = xerr.NewErrCode(xerr.FORM_PRASE_ERROR)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 视频流
func (l *FeedLogic) Feed(in *pb.FeedReq) (*pb.FeedResp, error) {

	var query = l.svcCtx.VideoModel.RowBuilder()
	if in.LastTime != nil {
		time, err := tool.UnixToTime(*(in.LastTime), "2006-01-02 15:04:05")
		if err != nil {
			return nil, errors.Wrapf(ErrDataFormatError, "字符串转时间错误 Unix:%+v,err:%+v", *(in.LastTime), err)
		}
		query = query.Where(squirrel.Gt{"create_time": time})
	}
	query = query.Limit(globalkey.FeedVideoNum)
	// TODO 加缓存
	videos, err := l.svcCtx.VideoModel.FindAll(l.ctx, query, "create_time  DESC")
	if err != nil {
		return nil, errors.Wrapf(ErrDBError, "user_id:%v,err:%v", *(in.UserId), err)
	}

	res := make([]*pb.Video, len(videos))
	_ = copier.Copy(&res, videos)

	if in.UserId != nil {
		getFavoriteInfoLogic := NewGetFavoriteInfoLogic(l.ctx, l.svcCtx)
		for _, video := range res {
			favoriteInfoResp, err := getFavoriteInfoLogic.GetFavoriteInfo(&pb.GetFavoriteInfoReq{
				UserId:  *in.UserId,
				VideoId: video.Id,
			})
			if err != nil {
				return nil, err
			}
			video.IsFavorite = favoriteInfoResp.IsFavorite
		}
	}

	return &pb.FeedResp{VideoList: res, NextTime: videos[0].CreateTime.Unix()}, nil
}
