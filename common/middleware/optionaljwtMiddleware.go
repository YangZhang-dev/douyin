package middleware

import (
	"context"
	"douyin/common/ctxdata"
	"douyin/common/result"
	"douyin/common/tool"
	"douyin/common/xerr"
	"net/http"
)

type OptionalJWTMiddleware struct {
	accessSecret string
}

func NewOptionalJWTMiddleware(accessSecret string) *OptionalJWTMiddleware {
	return &OptionalJWTMiddleware{accessSecret: accessSecret}
}

func (m *OptionalJWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_ = r.ParseForm()
		if r.Form.Has("token") {
			token := r.Form.Get("token")
			if token != "" {
				claims, err := tool.ParseToken(token, m.accessSecret)
				if err != nil {
					result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.TOKEN_PARSE_ERROR))
					return
				}
				ctx = context.WithValue(ctx, ctxdata.CtxKeyJwtUserId, claims.UserId)
			}
		}
		next(w, r.WithContext(ctx))
	}
}
