package ctxdata

import (
	"context"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	uid := ctx.Value(CtxKeyJwtUserId)
	if uid == nil {
		return 0
	}
	return uid.(int64)
}
