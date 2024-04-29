package models

import (
	"yim_server/commen/models"
	"yim_server/commen/models/ctype"
)

// UserConfModel 用户配置表
type UserConfModel struct {
	models.Model
	UserId         uint                  `json:"userId"`
	UserModel      UserModel             `json:"userModel" gorm:"foreignKey:UserId" json:"-"`
	RecallMsg      *string               `gorm:"size:32" json:"recallMsg"` //撤回消息的提示内容
	FriendOnline   bool                  `json:"friendOnline"`             //好友上线提醒
	Sound          bool                  `json:"sound"`                    //声音提醒
	SecureLink     bool                  `json:"secureLink"`               //安全链接
	SavePwd        bool                  `json:"savePwd"`                  //保存密码
	SearchUser     int8                  `json:"searchUser"`               //别人搜索用户的方式，0：不允许，1：允许用户号搜索，2：允许昵称搜索
	Verify         int8                  `json:"Verify"`                   //添加好友验证方式，0：不允许添加，1：不需要验证，2：需要验证，3：需要回答问题，4：需要正确回答问题
	VerifyQuestion *ctype.VerifyQuestion `json:"verifyQuestion"`           //好友验证问题 为3、4时需要
	IsOnline       bool                  `json:"isOnline"`                 //是否在线
}
