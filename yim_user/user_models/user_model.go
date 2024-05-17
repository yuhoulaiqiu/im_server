package user_models

import "yim_server/common/models"

// UserModel 用户表
type UserModel struct {
	models.Model
	Pwd            string         `gorm:"size:64" json:"-"`
	NickName       string         `gorm:"size:32" json:"nickName"`
	Abstract       string         `gorm:"size:128" json:"abstract"`
	Avatar         string         `gorm:"size:256" json:"avatar"`
	IP             string         `gorm:"size:32" json:"ips"`
	Addr           string         `gorm:"size:64" json:"addr"`
	OpenID         string         `gorm:"size:64" json:"-"`              //第三方登录的唯一标识
	RegisterSource string         `gorm:"size:32" json:"registerSource"` //注册来源
	Role           int8           `gorm:"size:4" json:"role"`            //1:管理员 2:普通用户
	UserConfModel  *UserConfModel `json:"userConfModel" gorm:"foreignKey:UserId" `
}
