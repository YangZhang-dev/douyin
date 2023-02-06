package chat

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

type HistoryMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryMessageLogic {
	return &HistoryMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryMessageLogic) HistoryMessage(req *types.HistoryMessageReq) (resp *types.HistoryMessageResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	getHistoryMessageResp, err := l.svcCtx.UserRpc.GetHistoryMessage(l.ctx, &pb.GetHistoryMessageReq{
		UserId:   userId,
		ToUserId: req.ToUserId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var res []types.Message
	messages := getHistoryMessageResp.MessageList
	_ = copier.Copy(&res, messages)

	return &types.HistoryMessageResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		Message: res,
	}, nil
}
