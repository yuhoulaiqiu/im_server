package chat_models

import (
	"yim_server/common/models"
	"yim_server/common/models/ctype"
)

type ChatModel struct {
	models.Model
	SendUserId uint            `json:"sendUserId"`                //发送者
	RecvUserId uint            `json:"recvUserId"`                //接收者
	MsgType    int8            `json:"msgType"`                   //消息类型 1：文本消息 2：图片消息 3：视频消息 4：文件消息 5：音频消息 6：语音通话 7：视频通话 8：撤回消息 9：回复消息 10：引用消息 11：位置消息
	MsgPreview string          `gorm:"size:64" json:"msgPreview"` //消息预览
	Msg        ctype.Msg       `json:"msg"`                       //消息内容
	SystemMsg  ctype.SystemMsg `json:"systemMsg"`                 //系统提示
}
