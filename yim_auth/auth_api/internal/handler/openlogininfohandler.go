package handler

import (
	"net/http"
	"yim_server/common/response"
	"yim_server/yim_auth/auth_api/internal/logic"
	"yim_server/yim_auth/auth_api/internal/svc"
)

func open_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOpen_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Open_login_info()
		response.Response(r, w, resp, err)
	}
}
