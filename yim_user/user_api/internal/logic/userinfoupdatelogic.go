package logic

import (
	"context"
	"errors"
	"yim_server/common/models/ctype"
	"yim_server/utils/maps"
	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"
	"yim_server/yim_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserInfoUpdate 是一个方法，它用于更新用户信息。
// 它接收一个 UserInfoUpdateRequest 类型的参数，并返回一个 UserInfoUpdateResponse 类型的响应和一个错误。
func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {
	// 使用反射将请求转换为 map
	userMaps := maps.ReflectToMap(*req, "user")
	// 如果 map 不为空
	if len(userMaps) != 0 {
		// 创建一个 UserModel 类型的变量
		var user user_models.UserModel
		// 从数据库中获取用户信息
		err := l.svcCtx.DB.Take(&user, req.UserID).Error
		// 如果获取失败，返回错误
		if err != nil {
			return nil, errors.New("用户不存在")
		}
		// 更新用户信息
		err = l.svcCtx.DB.Model(&user).Updates(userMaps).Error
		// 如果更新失败，返回错误
		if err != nil {
			logx.Errorf("更新失败:%v", err)
			return nil, errors.New("更新失败")
		}
	}
	// 使用反射将请求转换为 map
	userConfMaps := maps.ReflectToMap(*req, "user_conf")
	// 如果 map 不为空
	if len(userConfMaps) != 0 {
		// 创建一个 UserConfModel 类型的变量
		var userConf user_models.UserConfModel
		// 从数据库中获取用户配置信息
		err := l.svcCtx.DB.Take(&userConf, req.UserID).Error
		// 如果获取失败，返回错误
		if err != nil {
			return nil, errors.New("用户不存在")
		}
		// 如果 map 中包含 "verify_question" 键
		verifyQuestion, ok := userConfMaps["verify_question"]
		if ok {
			// 从 map 中删除 "verify_question" 键
			delete(userConfMaps, "verify_question")
			// 创建一个 VerifyQuestion 类型的变量
			data := ctype.VerifyQuestion{}
			// 将 map 转换为 VerifyQuestion 类型
			maps.MapToStruct(verifyQuestion.(map[string]interface{}), &data)
			// 更新用户配置信息中的验证问题
			l.svcCtx.DB.Model(&userConf).Updates(user_models.UserConfModel{
				VerifyQuestion: &data,
			})
		}
		// 更新用户配置信息
		err = l.svcCtx.DB.Model(&userConf).Updates(userConfMaps).Error
		// 如果更新失败，返回错误
		if err != nil {
			logx.Errorf("更新失败:%v", err)
			return nil, errors.New("更新失败")
		}
	}
	// 返回响应和错误
	return
}
