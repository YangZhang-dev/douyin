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

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCommentList 获取视频评论
func (l *GetCommentListLogic) GetCommentList(in *pb.GetCommentListReq) (*pb.GetCommentListResp, error) {

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

	var query = l.svcCtx.CommentModel.RowBuilder().Where(squirrel.Eq{"video_id": in.VideoId})
	// TODO 加缓存
	comments, err := l.svcCtx.CommentModel.FindAll(l.ctx, query, "create_time DESC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "videoId:%v,err:%v", in.VideoId, err)
	}
	var res []*pb.Comment
	_ = copier.Copy(&res, comments)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(comments); j++ {
			if res[i].Id == comments[j].Id {
				res[i].CreateTime = comments[j].CreateTime.Format("2006-01-02 15:04:05")
			}
		}
	}
	return &pb.GetCommentListResp{CommentList: res}, nil
}
