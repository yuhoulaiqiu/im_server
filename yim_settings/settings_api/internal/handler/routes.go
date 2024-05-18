// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"yim_server/yim_settings/settings_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/auth/open_login",
				Handler: open_login_infoHandler(serverCtx),
			},
		},
	)
}