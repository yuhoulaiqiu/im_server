package handler

import (
	"net/http"
	"yim_server/common/response"
	"yim_server/yim_user/user_api/internal/logic"
	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendNoteUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendNoteRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewFriendNoteUpdateLogic(r.Context(), svcCtx)
		resp, err := l.FriendNoteUpdate(&req)
		response.Response(r, w, resp, err)
	}
}
