package logic

import (
	"context"
	"errors"
	"fmt"
	"time"
	"yim_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
	"yim_server/yim_auth/auth_api/internal/svc"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp string, err error) {
	if token == "" {
		err = errors.New("token不能为空")
		return
	}
	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("token无效")
		return
	}
	now := time.Now()
	expiration := payload.ExpiresAt.Time.Sub(now)
	key := fmt.Sprintf("logout_%s", token)
	// 过期时间就是jwt的失效时间
	l.svcCtx.Redis.SetNX(key, "", expiration)
	resp = "退出成功"
	return resp, nil
}
