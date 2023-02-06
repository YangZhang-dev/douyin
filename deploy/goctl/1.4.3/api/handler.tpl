package {{.PkgName}}

import (
	"net/http"

    "douyin/common/result"
    "github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpResult(r,w,nil,xerr.NewErrMsg("参数错误"))
			return
		}

        if err := validator.New().StructCtx(r.Context(), req); err != nil{
            result.ParamErrorResult(r,w,err)
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		result.HttpResult(r, w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
	}
}
