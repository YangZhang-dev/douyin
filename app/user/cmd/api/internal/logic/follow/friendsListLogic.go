package follow

import (
	"context"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"douyin/app/user/cmd/api/internal/svc"
	"douyin/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendsListLogic {
	return &FriendsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendsListLogic) FriendsList(req *types.FriendsListReq) (resp *types.FriendsListResp, err error) {
	curUserid := ctxdata.GetUidFromCtx(l.ctx)
	getFriendsListResp, err := l.svcCtx.UserRpc.GetFriendsList(l.ctx, &pb.GetFriendsListReq{UserId: req.UserId, CurUserId: &curUserid})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	res := make([]types.User, len(getFriendsListResp.User))
	_ = copier.Copy(&res, getFriendsListResp.User)
	return &types.FriendsListResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		UserList: res,
	}, nil
}
