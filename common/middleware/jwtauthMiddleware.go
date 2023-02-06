package middleware

import (
	"context"
	"douyin/common/ctxdata"
	"douyin/common/result"
	"douyin/common/tool"
	"douyin/common/xerr"
	"github.com/pkg/errors"
	"net/http"
)

var ErrTokenParseError = xerr.NewErrCode(xerr.TOKEN_PARSE_ERROR)

type JwtAuthMiddleWare struct {
	accessSecret string
}

func NewJwtAuthMiddleWareMiddleware(accessSecret string) *JwtAuthMiddleWare {
	return &JwtAuthMiddleWare{accessSecret: accessSecret}
}

func (m *JwtAuthMiddleWare) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_ = r.ParseForm()

		var token string
		if r.Form.Has("token") {
			token = r.Form.Get("token")
		}
		if token != "" {
			claims, err := tool.ParseToken(token, m.accessSecret)
			if err != nil {
				result.HttpResult(r, w, nil, errors.Wrapf(ErrTokenParseError, "token错误 err:%+v", err))
				return
			}
			ctx = context.WithValue(ctx, ctxdata.CtxKeyJwtUserId, claims.UserId)
		} else {
			result.HttpResult(r, w, nil, errors.Wrapf(ErrTokenParseError, "token错误"))
			return
		}
		next(w, r.WithContext(ctx))
	}
}
