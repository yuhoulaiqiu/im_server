package models

import "yim_server/commen/models"

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserId     uint      `json:"sendUserId"` //发送者
	SendUserModel  UserModel `gorm:"foreignKey:SendUserId" json:"-"`
	RecvUserId     uint      `json:"recvUserId"` //接收者
	RecvUserModel  UserModel `gorm:"foreignKey:RecvUserId" json:"-"`
	Status         int8      `json:"status"`                        //状态 1：通过, 2：拒绝	, 3：忽略
	AdditionalMsg  string    `gorm:"size:128" json:"additionalMsg"` //附加消息
	VerifyQuestion string    `json:"verifyQuestion"`                //验证问题
}
