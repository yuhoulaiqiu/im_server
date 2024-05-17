package logic

import (
	"context"
	"errors"
	"fmt"
	"yim_server/utils/jwts"
	"yim_server/utils/open_login"
	"yim_server/yim_auth/auth_models"
	"yim_server/yim_user/user_rpc/types/user_rpc"

	"yim_server/yim_auth/auth_api/internal/svc"
	"yim_server/yim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {
	type OpenInfo struct {
		Ncikname string
		Avatar   string
		OpenID   string
	}
	var info OpenInfo
	switch req.Flag {
	case "qq":
		qqInfo, openErr := open_login.NewQQLogin(req.Code, open_login.QQConfig{
			AppID:    l.svcCtx.Config.QQ.AppID,
			AppKey:   l.svcCtx.Config.QQ.AppKey,
			Redirect: l.svcCtx.Config.QQ.Redirect,
		})
		info = OpenInfo{
			Ncikname: qqInfo.Nickname,
			Avatar:   qqInfo.Avatar,
			OpenID:   qqInfo.OpenID,
		}
		err = openErr
	default:
		err = errors.New("不支持的登录方式")
	}
	if err != nil {
		logx.Error(err)
		return nil, errors.New("qq登录失败")
	}
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		//注册逻辑
		fmt.Println("注册服务")
		res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
			Nickname:       info.Ncikname,
			Avatar:         info.Avatar,
			OpenId:         info.OpenID,
			Role:           2,
			Password:       "",
			RegisterSource: "qq",
		})
		if err != nil {
			logx.Error(err)
			return nil, errors.New("注册失败")
		}
		user.Model.ID = uint(res.UserId)
		user.Role = 2
		user.NickName = info.Ncikname
	}
	// 登录逻辑
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
