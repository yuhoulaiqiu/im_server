package handler

import (
	"net/http"
	"yim_server/common/response"
	"yim_server/yim_file/file_api/internal/logic"
	"yim_server/yim_file/file_api/internal/svc"
	"yim_server/yim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewImageShowLogic(r.Context(), svcCtx)
		err := l.ImageShow(&req)
		response.Response(r, w, nil, err)
	}
}
