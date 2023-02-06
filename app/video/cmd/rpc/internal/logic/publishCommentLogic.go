package logic

import (
	"context"
	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"
	"douyin/app/video/model"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishCommentLogic {
	return &PublishCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PublishComment 发布评论
func (l *PublishCommentLogic) PublishComment(in *pb.PublishCommentReq) (*pb.PublishCommentResp, error) {

	getVideoListByIdsLogic := NewGetVideoListByIdsLogic(l.ctx, l.svcCtx)
	getVideoListByIdsResp, err := getVideoListByIdsLogic.GetVideoListByIds(&pb.GetVideoListByIdsReq{
		UserId: nil,
		Ids:    []int64{in.VideoId},
	})
	if getVideoListByIdsResp.VideoList == nil {
		return nil, errors.Wrapf(ErrVideoNotExistError, "视频不存在 video_id:%+v", in.VideoId)
	}
	if err != nil {
		return nil, err
	}

	var video model.Video
	_ = copier.Copy(&video, getVideoListByIdsResp.VideoList[0])
	video.CommentCount++
	var lastId int64
	err = l.svcCtx.VideoModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		_, err = l.svcCtx.VideoModel.Update(l.ctx, nil, &video)
		if err != nil {
			return errors.Wrapf(ErrDBError, "videoId:%v,err:%v", in.VideoId, err)
		}

		lastIdRes, err := l.svcCtx.CommentModel.Insert(l.ctx, nil, &model.Comment{
			UserId:  in.UserId,
			VideoId: in.VideoId,
			Content: in.Content,
		})
		if err != nil {
			return errors.Wrapf(ErrDBError, "videoId:%v,err:%v", in.VideoId, err)
		}
		lastId, err = lastIdRes.LastInsertId()
		if err != nil {
			return errors.Wrapf(ErrDBError, "err:%v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, lastId)
	if err != nil {
		return nil, errors.Wrapf(ErrDBError, "commentId:%v,err:%v", lastId, err)
	}
	var res pb.Comment
	_ = copier.Copy(&res, comment)
	res.CreateTime = comment.CreateTime.Format("2006-01-02 15:04:05")
	return &pb.PublishCommentResp{Comment: &res}, nil
}
