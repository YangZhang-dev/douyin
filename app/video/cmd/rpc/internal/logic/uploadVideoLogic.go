package logic

import (
	"context"
	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"
	"douyin/app/video/model"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrDBError = xerr.NewErrCode(xerr.DB_ERROR)

type UploadVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadVideo 发布视频
func (l *UploadVideoLogic) UploadVideo(in *pb.UploadVideoReq) (*pb.UploadVideoResp, error) {

	_, err := l.svcCtx.VideoModel.Insert(l.ctx, nil, &model.Video{
		UserId:        in.UserId,
		Title:         in.Title,
		PlayUrl:       in.PlayUrl,
		CoverUrl:      in.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrDBError, "err:%v,userid:%+v,playurl:%+v", err, in.UserId, in.PlayUrl)
	}
	return &pb.UploadVideoResp{}, nil
}
