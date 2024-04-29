package models

import "yim_server/commen/models"

type UserModel struct {
	models.Model
	Pwd      string `json:"pwd"`
	NickName string `json:"nick_name"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
}
