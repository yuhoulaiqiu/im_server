package user_models

import (
	"gorm.io/gorm"
	"yim_server/common/models"
)

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserId    uint      `json:"sendUserId"`               //发送者A
	RecvUserId    uint      `json:"recvUserId"`               //接收者B
	SendNote      string    `gorm:"size:128" json:"sendNote"` //发送者备注 A->B的备注
	RecvNote      string    `gorm:"size:128" json:"recvNote"` //接收者备注 B->A的备注
	SendUserModel UserModel `gorm:"foreignKey:SendUserId" json:"-"`
	RecvUserModel UserModel `gorm:"foreignKey:RecvUserId" json:"-"`
}

func (f *FriendModel) IsFriend(db *gorm.DB, A, B uint) bool {
	err := db.Take(&f, "(send_user_id = ? and recv_user_id = ?) or (send_user_id = ? and recv_user_id = ?)", A, B, B, A).Error
	if err != nil {
		return false
	}
	return true
}

func (f *FriendModel) GetUserNote(userID uint) string {
	if userID == f.SendUserId {
		return f.SendNote
	} else {
		return f.RecvNote
	}
}
