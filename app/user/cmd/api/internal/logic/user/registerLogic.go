package user

import (
	"context"
	"douyin/app/user/cmd/api/internal/svc"
	"douyin/app/user/cmd/api/internal/types"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {

	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.RegisterResp{
		UserId: registerResp.UserId,
		Token:  registerResp.Token,
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
	}, nil
}
