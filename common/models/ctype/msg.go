package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Msg struct {
	Type         int8          `json:"type"`         //消息类型
	Content      *string       `json:"content"`      //为1时需要
	ImgMsg       *ImgMsg       `json:"imgMsg"`       //为2时需要
	VideoMsg     *VideoMsg     `json:"videoMsg"`     //为3时需要
	FileMsg      *FileMsg      `json:"fileMsg"`      //为4时需要
	VoiceMsg     *VoiceMsg     `json:"voiceMsg"`     //为5时需要
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg"` //为6时需要
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg"` //为7时需要
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg"`  //为8时需要
	ReplyMsg     *ReplyMsg     `json:"replyMsg"`     //为9时需要
	QuoteMsg     *QuoteMsg     `json:"quoteMsg"`     //为10时需要
	LocationMsg  *LocationMsg  `json:"locationMsg"`  //为11时需要
	AtMsg        *AtMsg        `json:"atMsg"`        //为12时需要
}

func (c *Msg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}
func (c *Msg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type ImgMsg struct {
	Title string `json:"title"` //图片标题
	Src   string `json:"src"`   //图片地址
}
type VideoMsg struct {
	Title string `json:"title"` //视频标题
	Src   string `json:"src"`   //视频地址
	Time  int    `json:"time"`  //视频时长,单位秒
}
type FileMsg struct {
	Title string `json:"title"` //文件标题
	Src   string `json:"src"`   //文件地址
	Size  int    `json:"size"`  //文件大小,单位字节
	Type  string `json:"type"`  //文件类型
}
type VoiceMsg struct {
	Src  string `json:"src"`  //音频地址
	Time int    `json:"time"` //音频时长,单位秒
}
type VoiceCallMsg struct {
	StartTime time.Time `json:"startTime"` //通话开始时间
	EndTime   time.Time `json:"endTime"`   //通话结束时间
	EndReason int8      `json:"endReason"` //通话结束原因,0:发起方挂断,1:接收方挂断,2:对方未接听,3:对方拒绝接听,4:对方忙线中,5:网络错误,6:对方不在线,8:已取消
}
type VideoCallMsg struct {
	StartTime time.Time `json:"startTime"` //通话开始时间
	EndTime   time.Time `json:"endTime"`   //通话结束时间
	EndReason int8      `json:"endReason"` //通话结束原因,0:发起方挂断,1:接收方挂断,2:对方未接听,3:对方拒绝接听,4:对方忙线中,5:网络错误,6:对方不在线,8:已取消
}
type WithdrawMsg struct {
	Content   string `json:"content"` //撤回消息的提示词
	OriginMsg *Msg   `json:"-"`       //被撤回的消息
}

// ReplyMsg 回复消息,目前只能回复文本消息
type ReplyMsg struct {
	MsgID   uint   `json:"msgID"`   //回复的消息ID
	Content string `json:"content"` //回复的内容
	Msg     *Msg   `json:"-"`       //回复的消息
}
type QuoteMsg struct {
	MsgID   uint   `json:"msgID"`   //回复的消息ID
	Content string `json:"content"` //回复的内容
	Msg     *Msg   `json:"-"`       //回复的消息
}
type LocationMsg struct {
	Latitude  float64 `json:"latitude"`  //纬度
	Longitude float64 `json:"longitude"` //经度
	Address   string  `json:"address"`   //地址
}
type AtMsg struct {
	UserId  uint   `json:"userId"`  //被@的用户ID
	Content string `json:"content"` //回复的文本消息
	Msg     *Msg   `json:"-"`       //被@的消息
}
