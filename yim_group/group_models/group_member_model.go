package models3

import "yim_server/commen/models"

type GroupMemberModel struct {
	models.Model
	GroupId        uint       `json:"groupId"`
	GroupModel     GroupModel `gorm:"foreignKey:GroupId" json:"-"`
	UserId         uint       `json:"userId"`                        //用户ID
	Role           int8       `json:"role"`                          //角色 1：群主, 2：管理员, 3：普通成员
	MemberNickname string     `gorm:"size:32" json:"memberNickname"` //群昵称
	ProhibitTime   *int64     `json:"prohibitTime"`                  //禁言时间，单位分
}
