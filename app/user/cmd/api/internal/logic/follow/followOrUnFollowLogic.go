package follow

import (
	"context"
	"douyin/app/user/cmd/api/internal/svc"
	"douyin/app/user/cmd/api/internal/types"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/globalkey"
	"douyin/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserCannotActionSelfError = xerr.NewErrCode(xerr.USER_CANNOT_ACTION_SELF)
var ErrRequestParamError = xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)

type FollowOrUnFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowOrUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowOrUnFollowLogic {
	return &FollowOrUnFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowOrUnFollowLogic) FollowOrUnFollow(req *types.FollowOrUnfollowReq) (resp *types.FollowOrUnfollowResp, err error) {
	curUserId := ctxdata.GetUidFromCtx(l.ctx)
	if curUserId == req.ToUserId {
		return nil, errors.Wrapf(ErrUserCannotActionSelfError, "req: %+v", req)
	}

	if req.ActionType == globalkey.Follow {
		_, err = l.svcCtx.UserRpc.Follow(l.ctx, &pb.FollowReq{
			UserId:   curUserId,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
	} else if req.ActionType == globalkey.Unfollow {
		_, err = l.svcCtx.UserRpc.UnFollow(l.ctx, &pb.UnFollowReq{
			UserId:   curUserId,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
	} else {
		return nil, errors.Wrapf(ErrRequestParamError, "请求参数错误,action_type:%+v", req.ActionType)
	}

	return &types.FollowOrUnfollowResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
	}, nil
}
