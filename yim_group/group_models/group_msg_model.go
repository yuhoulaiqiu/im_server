package models3

import (
	"yim_server/commen/models"
	"yim_server/commen/models/ctype"
)

// GroupMsgModel 群消息表
type GroupMsgModel struct {
	models.Model
	GroupId    uint            `json:"groupId"` //群ID
	GroupModel GroupModel      `gorm:"foreignKey:GroupId" json:"-"`
	SendUserId uint            `json:"sendUserId"`           //发送者
	MsgPreview string          `gorm:"64" json:"msgPreview"` //消息预览
	Msg        ctype.Msg       `json:"msg"`                  //消息内容
	MsgType    int8            `json:"msgType"`              //消息类型 1：文本消息 2：图片消息 3：视频消息 4：文件消息 5：音频消息 6：语音通话 7：视频通话 8：撤回消息 9：回复消息 10：引用消息 11：位置消息 12:@消息
	SystemMsg  ctype.SystemMsg `json:"systemMsg"`            //系统提示
}
