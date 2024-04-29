package models

import "yim_server/commen/models"

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserId    uint      `json:"sendUserId"`           //发送者
	RecvUserId    uint      `json:"recvUserId"`           //接收者
	Note          string    `gorm:"size:128" json:"note"` //备注
	SendUserModel UserModel `gorm:"foreignKey:SendUserId" json:"-"`
	RecvUserModel UserModel `gorm:"foreignKey:RecvUserId" json:"-"`
}
