package models3

import "yim_server/commen/models"

// GroupVerifyModel 群验证信息表
type GroupVerifyModel struct {
	models.Model
	GroupId        uint       `json:"groupId"` //群ID
	GroupModel     GroupModel `gorm:"foreignKey:GroupId" json:"-"`
	UserId         uint       `json:"userId"`                       //用户ID
	Status         int8       `json:"status"`                       //状态 1：通过, 2：拒绝	, 3：忽略
	AdditionalMsg  string     `gorm:"size:32" json:"additionalMsg"` //附加消息
	VerifyQuestion string     `json:"verifyQuestion"`               //验证问题
	Type           int8       `json:"type"`                         //类型 1：申请加群, 2：退群
}
