package follow

import (
	"douyin/common/xerr"
	"net/http"

	"douyin/app/user/cmd/api/internal/logic/follow"
	"douyin/app/user/cmd/api/internal/svc"
	"douyin/app/user/cmd/api/internal/types"
	"douyin/common/result"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, xerr.NewErrMsg("参数错误"))
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, xerr.NewErrMsg("参数错误"))
			return
		}

		l := follow.NewFollowListLogic(r.Context(), svcCtx)
		resp, err := l.FollowList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
