package logic

import (
	"context"

	"yim_server/yim_auth/auth_api/internal/svc"
	"yim_server/yim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_login_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_login_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_login_infoLogic {
	return &Open_login_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_login_infoLogic) OpenLoginInfo() (resp []types.OpenLoginResponse, err error) {
	for _, s := range l.svcCtx.Config.OpenLoginList {
		resp = append(resp, types.OpenLoginResponse{
			Name: s.Name,
			Icon: s.Icon,
			Href: s.Href,
		})
	}
	return

}
