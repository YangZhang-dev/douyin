package favorite

import (
	"context"
	userpb "douyin/app/user/cmd/rpc/pb"
	"douyin/app/video/cmd/api/internal/logic/video"
	"douyin/app/video/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNotExistError = xerr.NewErrCode(xerr.USER_NOT_EXIST_ERROR)

type FavoriteVideosListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteVideosListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideosListLogic {
	return &FavoriteVideosListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteVideosListLogic) FavoriteVideosList(req *types.FavoriteVideosListReq) (resp *types.FavoriteVideosListResp, err error) {
	// 1. 校验数据
	userId := ctxdata.GetUidFromCtx(l.ctx)
	var curUserId *int64
	if userId != 0 {
		curUserId = &userId
	}
	findUserResp, err := l.svcCtx.UserRpc.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
		UserId: nil,
		Ids:    []int64{req.UserId},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	if findUserResp.Users == nil {
		return nil, errors.Wrapf(ErrUserNotExistError, "用户不存在")
	}

	// 2. 获取点赞视频列表
	favoriteListResp, err := l.svcCtx.VideoRpc.FavoriteList(l.ctx, &pb.FavoriteListReq{
		UserId:    req.UserId,
		CurUserId: curUserId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	// 3. 获取作者信息
	videos := favoriteListResp.VideoList
	authorsId, res := video.DrawAuthorId(videos)

	userListByIdsResp, err := l.svcCtx.UserRpc.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
		UserId: curUserId,
		Ids:    authorsId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	// 4. 将作者信息封装到视频信息中
	authors := userListByIdsResp.Users
	video.SetAuthorInfo(res, authors)

	return &types.FavoriteVideosListResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		VideoList: res,
	}, nil
}
