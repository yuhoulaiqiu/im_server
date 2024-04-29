package models

import "yim_server/commen/models"

type FriendModel struct {
	models.Model
	SendUserId uint   `json:"send_user_id"` //发送者
	RecvUserId uint   `json:"recv_user_id"` //接收者
	Note       string `json:"note"`         //备注
	
}
