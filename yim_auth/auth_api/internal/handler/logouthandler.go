package handler

import (
	"net/http"
	"yim_server/common/response"
	"yim_server/yim_auth/auth_api/internal/logic"
	"yim_server/yim_auth/auth_api/internal/svc"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		resp, err := l.Logout(token)
		response.Response(r, w, resp, err)
	}
}
