package favorite

import (
	"douyin/common/xerr"
	"net/http"

	"douyin/app/video/cmd/api/internal/logic/favorite"
	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"
	"douyin/common/result"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteOrUnfavoriteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteOrUnfavoriteReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, xerr.NewErrMsg("参数错误"))
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, xerr.NewErrMsg("参数错误"))
			return
		}

		l := favorite.NewFavoriteOrUnfavoriteLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteOrUnfavorite(&req)
		result.HttpResult(r, w, resp, err)
	}
}