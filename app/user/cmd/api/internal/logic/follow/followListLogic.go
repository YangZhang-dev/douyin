package follow

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

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListReq) (resp *types.FollowListResp, err error) {
	curUserId := ctxdata.GetUidFromCtx(l.ctx)
	getFollowListByUserIdResp, err := l.svcCtx.UserRpc.GetFollowListByUserId(l.ctx, &pb.GetFollowListByUserIdReq{UserId: req.UserId, CurUserId: &curUserId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	users := make([]types.User, len(getFollowListByUserIdResp.Users))
	_ = copier.Copy(&users, getFollowListByUserIdResp.Users)

	return &types.FollowListResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		UserList: users,
	}, nil
}
