package logic

import (
	"context"
	"douyin/app/user/model"
	"douyin/common/tool"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserOrPasswordError = xerr.NewErrCode(xerr.USER_OR_PASSWORD_ERROR)
var ErrTokenGenerateError = xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR)

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 1. 检查用户和密码的正确性
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "username:%s,err:%v", in.Username, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserOrPasswordError, "用户不存在 username:%s,err:%v", in.Username, err)
	}
	isCorrect := tool.ComparePasswords(user.Password, in.Password)
	if !isCorrect {
		return nil, errors.Wrapf(ErrUserOrPasswordError, "密码错误 password:%s", in.Password)
	}

	// 2. 生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{UserId: user.Id})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		UserId: user.Id,
		Token:  token.Token,
	}, nil
}
