package video

import (
	"douyin/common/globalkey"
	"douyin/common/xerr"
	"github.com/pkg/errors"
	"net/http"

	"douyin/app/video/cmd/api/internal/logic/video"
	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"
	"douyin/common/result"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var ErrUploadFileOverSize = xerr.NewErrCode(xerr.UPLOAD_FILE_LIMIT_EXCEEDED)
var ErrUploadFileNotFound = xerr.NewErrCode(xerr.UPLOAD_FILE_NOT_FOUND_ERROR)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if r.MultipartForm.File["data"] == nil {
			result.HttpResult(r, w, nil, ErrUploadFileNotFound)
			return
		}
		file := r.MultipartForm.File["data"][0]
		if file.Size > globalkey.MaxVideoSize {
			result.HttpResult(r, w, nil, errors.Wrapf(ErrUploadFileOverSize, "上传文件大小超出限制 target:video/mp4, size:%+v", file.Size))
			return
		}

		l := video.NewPublishVideoLogic(r.Context(), svcCtx, file)
		resp, err := l.PublishVideo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
