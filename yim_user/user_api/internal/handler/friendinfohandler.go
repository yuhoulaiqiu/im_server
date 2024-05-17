package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"yim_server/common/response"
	"yim_server/yim_user/user_api/internal/logic"
	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("httpx.Parse(r, &req) error(%v)", err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewFriendInfoLogic(r.Context(), svcCtx)
		resp, err := l.FriendInfo(&req)
		response.Response(r, w, resp, err)
	}
}
