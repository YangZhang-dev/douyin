package logic

import (
	"context"
	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/app/user/model"
	"douyin/common/tool"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyExistError = xerr.NewErrCode(xerr.USER_ALREADY_EXIST_ERROR)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// user
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 1. 检查用户名是否存在
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "username:%s,err:%v", in.Username, err)
	}
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyExistError, "用户已经存在 username:%s,err:%v", in.Username, err)
	}

	// 2. 加密密码
	pwdHash, err := tool.HashAndSalt(in.Password)
	if err != nil {
		return nil, errors.Wrapf(ErrDataFormatError, "数据校验错误 password:%s,err:%v", in.Password, err)
	}

	// 3. 创建用户
	result, err := l.svcCtx.UserModel.Insert(l.ctx, nil, &model.User{
		Username:      in.Username,
		Password:      pwdHash,
		FollowCount:   0,
		FollowerCount: 0,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrDBError, "err:%v,user:%+v", err, user)
	}

	// 4. 获取创建用户id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(ErrDBError, "insertResult.LastInsertId err:%v,user:%+v", err, user)
	}

	// 5. 生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{UserId: id})
	if err != nil {
		return nil, errors.Wrapf(ErrTokenGenerateError, "生成token错误:%v,id:%+v", err, id)
	}

	return &pb.RegisterResp{
		UserId: id,
		Token:  token.Token,
	}, nil
}
