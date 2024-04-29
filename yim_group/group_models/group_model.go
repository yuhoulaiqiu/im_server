package models3

import (
	"yim_server/commen/models"
	"yim_server/commen/models/ctype"
)

type GroupModel struct {
	models.Model
	Title              string                `gorm:"size:32" json:"title"`     // 群名称
	Avatar             string                `gorm:"size:128" json:"avatar"`   // 群头像
	Abstract           string                `gorm:"size:256" json:"abstract"` // 群简介
	Creator            uint                  `json:"creator"`                  // 群主
	Size               int                   `json:"size"`                     // 群容量 20 100 200 500 1000 2000
	IsSearch           bool                  `json:"isSearch"`                 // 是否可以被搜索到
	Verify             int8                  `json:"Verify"`                   //加群验证方式，0：不允许添加，1：不需要验证，2：需要验证，3：需要回答问题，4：需要正确回答问题
	VerifyQuestion     *ctype.VerifyQuestion `json:"verifyQuestion"`           //验证问题 为3、4时需要
	IsInvite           bool                  `json:"isInvite"`                 // 是否允许群成员邀请
	IsTemporarySession bool                  `json:"IsTemporarySession"`       // 是否允许临时会话
	IsProhibit         bool                  `json:"isProhibit"`               // 是否开启全员禁言
}
