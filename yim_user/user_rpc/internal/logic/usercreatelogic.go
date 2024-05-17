package logic

import (
	"context"
	"errors"
	"yim_server/yim_user/user_models"

	"yim_server/yim_user/user_rpc/internal/svc"
	"yim_server/yim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	var user user_models.UserModel
	err := l.svcCtx.DB.Take(&user, "open_id = ?", in.OpenId).Error
	if err == nil {
		return nil, errors.New("用户已存在")
	}
	user = user_models.UserModel{
		OpenID:         in.OpenId,
		Role:           int8(in.Role),
		NickName:       in.Nickname,
		Avatar:         in.Avatar,
		RegisterSource: in.RegisterSource,
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建用户失败")
	}
	// 创建用户配置
	l.svcCtx.DB.Create(&user_models.UserConfModel{
		UserId:       user.ID,
		RecallMsg:    nil,   //撤回消息的提示内容
		FriendOnline: false, //关闭好友上线提醒
		Sound:        true,  //开启声音提醒
		SecureLink:   false, //关闭安全链接
		SavePwd:      false, //关闭保存密码
		SearchUser:   2,
		Verify:       2,
		IsOnline:     true,
	})
	return &user_rpc.UserCreateResponse{UserId: int32(user.ID)}, nil
}
