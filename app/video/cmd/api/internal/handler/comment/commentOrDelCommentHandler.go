package comment

import (
	"douyin/common/xerr"
	"net/http"

	"douyin/app/video/cmd/api/internal/logic/comment"
	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"
	"douyin/common/result"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentOrDelCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentOrDelCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, xerr.NewErrMsg("参数错误"))
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, xerr.NewErrMsg("参数错误"))
			return
		}

		l := comment.NewCommentOrDelCommentLogic(r.Context(), svcCtx)
		resp, err := l.CommentOrDelComment(&req)
		result.HttpResult(r, w, resp, err)
	}
}
