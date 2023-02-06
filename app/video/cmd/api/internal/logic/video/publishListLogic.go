package video

import (
	"context"
	userpb "douyin/app/user/cmd/rpc/pb"
	"douyin/app/video/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNotExistError = xerr.NewErrCode(xerr.USER_NOT_EXIST_ERROR)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListReq) (resp *types.PublishListResp, err error) {
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

	publishListResp, err := l.svcCtx.VideoRpc.PublishList(l.ctx, &pb.PublishListReq{UserId: req.UserId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	// 3. 获取作者信息
	videos := publishListResp.VideoList
	authorsId, res := DrawAuthorId(videos)

	userListByIdsResp, err := l.svcCtx.UserRpc.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
		UserId: curUserId,
		Ids:    authorsId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	// 4. 将作者信息封装到视频信息中
	authors := userListByIdsResp.Users
	SetAuthorInfo(res, authors)

	return &types.PublishListResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		VideoList: res,
	}, nil
}
