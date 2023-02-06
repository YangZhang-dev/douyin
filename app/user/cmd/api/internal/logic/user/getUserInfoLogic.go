package user

import (
	"context"
	"douyin/app/user/cmd/api/internal/svc"
	"douyin/app/user/cmd/api/internal/types"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	curUserId := ctxdata.GetUidFromCtx(l.ctx)
	var isFollow = false
	// 1. 获取用户信息
	getUserByIdResp, err := l.svcCtx.UserRpc.GetUserListByIds(l.ctx, &pb.GetUserListByIdsReq{
		Ids: []int64{req.UserId},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	// 2. 如果登录了，再获取关注状态
	if curUserId != 0 {
		getFollowInfoResp, err := l.svcCtx.UserRpc.GetFollowInfo(l.ctx, &pb.GetFollowInfoReq{
			UserId:   curUserId,
			ToUserId: req.UserId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		isFollow = getFollowInfoResp.IsFollow
	}
	// 3.组装数据并返回
	var user types.User
	_ = copier.Copy(&user, getUserByIdResp.Users[0])
	user.IsFollow = isFollow

	return &types.UserInfoResp{
		User: user,
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
	}, nil
}
