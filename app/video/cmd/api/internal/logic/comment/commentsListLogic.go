package comment

import (
	"context"
	userpb "douyin/app/user/cmd/rpc/pb"
	videopb "douyin/app/video/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentsListLogic {
	return &CommentsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentsListLogic) CommentsList(req *types.CommentsListReq) (resp *types.CommentsListResp, err error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)

	getCommentListResp, err := l.svcCtx.VideoRpc.GetCommentList(l.ctx, &videopb.GetCommentListReq{VideoId: req.VideoId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	comments := getCommentListResp.CommentList
	var res []types.Comment
	_ = copier.Copy(&res, comments)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(comments); j++ {
			if res[i].Id == comments[j].Id {
				res[i].CreateDate = comments[j].CreateTime
			}
		}
	}

	var ids []int64
	for i := 0; i < len(comments); i++ {
		ids = append(ids, comments[i].UserId)
		for j := 0; j < len(res); j++ {
			if res[j].Id == comments[i].Id {
				res[j].User.Id = comments[i].UserId
			}
		}
	}
	getUserListByIdsResp, err := l.svcCtx.UserRpc.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
		UserId: &userId,
		Ids:    ids,
	})
	if err != nil {
		return nil, err
	}
	users := getUserListByIdsResp.Users

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(users); j++ {
			if res[i].User.Id == users[j].Id {
				_ = copier.Copy(&res[i].User, users[j])
			}
		}
	}
	return &types.CommentsListResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		CommentsList: res,
	}, nil
}
