package middleware

import (
	"context"
	"douyin/common/ctxdata"
	"douyin/common/globalkey"
	"douyin/common/result"
	"douyin/common/tool"
	"douyin/common/xerr"
	"github.com/pkg/errors"
	"net/http"
)

var ErrFormParseError = xerr.NewErrCode(xerr.FORM_PRASE_ERROR)
var ErrTokenParseError = xerr.NewErrCode(xerr.TOKEN_PARSE_ERROR)

type ParseFormMiddleware struct {
	accessSecret string
}

func NewParseFormMiddleware(accessSecret string) *ParseFormMiddleware {
	return &ParseFormMiddleware{accessSecret: accessSecret}
}

func (m *ParseFormMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := r.ParseMultipartForm(globalkey.MaxVideoSize)
		if err != nil {
			result.HttpResult(r, w, nil, errors.Wrapf(ErrFormParseError, "表单解析错误 err:%+v", err))
			return
		} else if r.MultipartForm.Value["token"] != nil {
			claims, err := tool.ParseToken(r.MultipartForm.Value["token"][0], m.accessSecret)
			if err != nil {
				result.HttpResult(r, w, nil, errors.Wrapf(ErrTokenParseError, "token错误 err:%+v", err))
				return
			}
			ctx = context.WithValue(ctx, ctxdata.CtxKeyJwtUserId, claims.UserId)
		} else {
			result.HttpResult(r, w, nil, errors.Wrapf(ErrTokenParseError, "token错误 err:%+v", err))
			return
		}

		next(w, r.WithContext(ctx))
	}
}
