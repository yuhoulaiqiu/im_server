package logic

import (
	"context"
	"errors"
	"yim_server/utils/jwts"
	"yim_server/utils/pwd"
	"yim_server/yim_auth/auth_models"

	"yim_server/yim_auth/auth_api/internal/svc"
	"yim_server/yim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "id = ?", req.UserName).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		return nil, err
	}
	if !pwd.CheckPwd(user.Pwd, req.Password) {
		err = errors.New("用户名或密码错误")
		return nil, err
	}
	token, err := jwts.GenToken(jwts.JwtPayload{
		UserID:   user.ID,
		Nickname: user.NickName,
		Role:     int(user.Role),
	}, l.svcCtx.Config.Auth.AccessSecret, int(l.svcCtx.Config.Auth.AccessExpire))
	if err != nil {
		logx.Error(err)
		err = errors.New("生成token失败")
		return nil, err
	}
	return &types.LoginResponse{Token: token}, nil
}
