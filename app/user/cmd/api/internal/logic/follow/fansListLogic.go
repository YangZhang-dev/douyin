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

type FansListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansListLogic {
	return &FansListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FansListLogic) FansList(req *types.FansListReq) (resp *types.FansListResp, err error) {
	curUserId := ctxdata.GetUidFromCtx(l.ctx)
	getFollowListByUserIdResp, err := l.svcCtx.UserRpc.GetFansListByUserId(l.ctx, &pb.GetFansListByUserIdReq{UserId: req.UserId, CurUserId: &curUserId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	users := make([]types.User, len(getFollowListByUserIdResp.Users))
	_ = copier.Copy(&users, getFollowListByUserIdResp.Users)

	return &types.FansListResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		UserList: users,
	}, nil

}
