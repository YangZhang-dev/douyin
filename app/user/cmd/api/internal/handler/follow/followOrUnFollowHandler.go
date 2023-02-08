package follow

import (
	"net/http"

	"douyin/app/user/cmd/api/internal/logic/follow"
	"douyin/app/user/cmd/api/internal/svc"
	"douyin/app/user/cmd/api/internal/types"
	"douyin/common/result"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowOrUnFollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowOrUnfollowReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := follow.NewFollowOrUnFollowLogic(r.Context(), svcCtx)
		resp, err := l.FollowOrUnFollow(&req)
		result.HttpResult(r, w, resp, err)
	}
}
