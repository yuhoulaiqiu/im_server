package handler

import (
	"net/http"
	"yim_server/common/response"
	"yim_server/yim_auth/auth_api/internal/logic"
	"yim_server/yim_auth/auth_api/internal/svc"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication()
		response.Response(r, w, resp, err)
	}
}
