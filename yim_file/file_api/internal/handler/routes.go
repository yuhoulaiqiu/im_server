// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"yim_server/yim_file/file_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/file/image",
				Handler: ImageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/file/uploads/:imagesType/:imageName",
				Handler: ImageShowHandler(serverCtx),
			},
		},
	)
}
