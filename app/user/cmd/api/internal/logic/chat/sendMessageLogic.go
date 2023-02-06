package chat

import (
	"context"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"douyin/app/user/cmd/api/internal/svc"
	"douyin/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageReq) (resp *types.SendMessageResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.UserRpc.SendMessage(l.ctx, &pb.SendMessageReq{
		UserId:   userId,
		ToUserId: req.ToUserId,
		Content:  req.Content,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.SendMessageResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
	}, nil
}
