package comment

import (
	"net/http"

	"douyin/app/video/cmd/api/internal/logic/comment"
	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"
	"douyin/common/result"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentsListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewCommentsListLogic(r.Context(), svcCtx)
		resp, err := l.CommentsList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
