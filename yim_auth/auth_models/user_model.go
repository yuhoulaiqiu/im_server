package models

import "yim_server/commen/models"

// UserModel 用户表
type UserModel struct {
	models.Model
	Pwd      string `gorm:"size:64" json:"pwd"`
	NickName string `gorm:"size:32" json:"nickName"`
	Abstract string `gorm:"size:128" json:"abstract"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	IP       string `gorm:"size:32" json:"ip"`
	Addr     string `gorm:"size:64" json:"addr"`
}
