package logic

import (
	"context"
	"douyin/app/video/cmd/rpc/internal/svc"
	"douyin/app/video/cmd/rpc/pb"
	"douyin/app/video/model"
	"douyin/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrCommentNotExistError = xerr.NewErrCode(xerr.COMMENT_NOT_EXIST_ERROR)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteComment 删除评论
func (l *DeleteCommentLogic) DeleteComment(in *pb.DeleteCommentReq) (*pb.DeleteCommentResp, error) {
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.CommentId)
	if comment == nil {
		return nil, errors.Wrapf(ErrCommentNotExistError, "评论不存在 comment_id:%+v", in.CommentId)
	}
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "commentId:%v,err:%v", in.CommentId, err)
	}
	getVideoListByIdsLogic := NewGetVideoListByIdsLogic(l.ctx, l.svcCtx)
	getVideoListByIdsResp, err := getVideoListByIdsLogic.GetVideoListByIds(&pb.GetVideoListByIdsReq{
		UserId: nil,
		Ids:    []int64{in.VideoId},
	})
	if err != nil {
		return nil, err
	}
	if getVideoListByIdsResp.VideoList == nil {
		return nil, errors.Wrapf(ErrVideoNotExistError, "视频不存在 video_id:%+v", in.VideoId)
	}
	var video model.Video
	_ = copier.Copy(&video, getVideoListByIdsResp.VideoList[0])
	video.CommentCount--

	err = l.svcCtx.VideoModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		_, err = l.svcCtx.VideoModel.Update(l.ctx, nil, &video)
		if err != nil {
			return errors.Wrapf(ErrDBError, "videoId:%v,err:%v", in.VideoId, err)
		}

		err := l.svcCtx.CommentModel.Delete(l.ctx, session, in.CommentId)
		if err != nil {
			return errors.Wrapf(ErrDBError, "commentId:%v,err:%v", in.CommentId, err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteCommentResp{}, nil
}
