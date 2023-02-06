package favorite

import (
	"context"
	"douyin/app/video/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/globalkey"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrRequestParamError = xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)

type FavoriteOrUnfavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteOrUnfavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteOrUnfavoriteLogic {
	return &FavoriteOrUnfavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteOrUnfavoriteLogic) FavoriteOrUnfavorite(req *types.FavoriteOrUnfavoriteReq) (resp *types.FavoriteOrUnfavoriteResp, err error) {

	curUserId := ctxdata.GetUidFromCtx(l.ctx)

	if req.ActionType == globalkey.Favorite {
		_, err = l.svcCtx.VideoRpc.Favorite(l.ctx, &pb.FavoriteReq{
			UserId:  curUserId,
			VideoId: req.VideoId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
	} else if req.ActionType == globalkey.UnFavorite {
		_, err = l.svcCtx.VideoRpc.UnFavorite(l.ctx, &pb.UnFavoriteReq{
			UserId:  curUserId,
			VideoId: req.VideoId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
	} else {
		return nil, errors.Wrapf(ErrRequestParamError, "请求参数错误,action_type:%+v", req.ActionType)
	}

	return &types.FavoriteOrUnfavoriteResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
	}, nil
}
