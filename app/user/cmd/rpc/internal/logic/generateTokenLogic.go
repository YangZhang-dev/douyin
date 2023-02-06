package logic

import (
	"context"
	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}
type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {

	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(ErrTokenGenerateError, "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}

	return &pb.GenerateTokenResp{
		Token: accessToken,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, seconds, userId int64) (string, error) {
	iat := time.Now()
	claim := MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(iat.Add(time.Second * time.Duration(seconds))),
			IssuedAt:  jwt.NewNumericDate(iat),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	return token.SignedString([]byte(secretKey))
}
