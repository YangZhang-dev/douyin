package comment

import (
	"context"
	userpb "douyin/app/user/cmd/rpc/pb"
	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"
	videopb "douyin/app/video/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/globalkey"
	"douyin/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrRequestParamError = xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)

type CommentOrDelCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentOrDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentOrDelCommentLogic {
	return &CommentOrDelCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentOrDelCommentLogic) CommentOrDelComment(req *types.CommentOrDelCommentReq) (resp *types.CommentOrDelCommentResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	if req.ActionType == globalkey.PublishComment {
		commentResp, err := l.svcCtx.VideoRpc.PublishComment(l.ctx, &videopb.PublishCommentReq{
			VideoId: req.VideoId,
			Content: req.CommentText,
			UserId:  userId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		getUserListByIdsResp, err := l.svcCtx.UserRpc.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
			UserId: &userId,
			Ids:    []int64{commentResp.Comment.UserId},
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		var res types.Comment
		_ = copier.Copy(&res, commentResp.Comment)
		_ = copier.Copy(&res.User, getUserListByIdsResp.Users[0])
		res.CreateDate = commentResp.Comment.CreateTime
		return &types.CommentOrDelCommentResp{
			Status: types.Status{
				StatusCode: xerr.OK,
				StatusMsg:  xerr.MapErrMsg(xerr.OK),
			},
			Comment: res,
		}, nil
	} else if req.ActionType == globalkey.DeleteComment {
		_, err := l.svcCtx.VideoRpc.DeleteComment(l.ctx, &videopb.DeleteCommentReq{
			VideoId:   req.VideoId,
			CommentId: req.CommentId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		return &types.CommentOrDelCommentResp{
			Status: types.Status{
				StatusCode: xerr.OK,
				StatusMsg:  xerr.MapErrMsg(xerr.OK),
			},
			Comment: types.Comment{},
		}, nil
	} else {
		return nil, errors.Wrapf(ErrRequestParamError, "请求参数错误,action_type:%+v", req.ActionType)
	}

}
