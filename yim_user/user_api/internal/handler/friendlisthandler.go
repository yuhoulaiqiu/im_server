package handler

import (
	"net/http"
	"yim_server/common/response"
	"yim_server/yim_user/user_api/internal/logic"
	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewFriendListLogic(r.Context(), svcCtx)
		resp, err := l.FriendList(&req)
		response.Response(r, w, resp, err)
	}
}
