// Code generated by goctl. DO NOT EDIT.
package types

type FriendInfoRequest struct {
	UserID   uint `header:"User-ID"`
	Role     int8 `header:"Role"`
	FriendID uint `form:"friendId"`
}

type FriendInfoResponse struct {
	UserID   uint   `json:"userId"`
	NickName string `json:"nickName"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
	Note     string `json:"note"` //备注
}

type FriendListRequest struct {
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
	UserID uint `header:"User-ID"`
	Role   int8 `header:"Role"`
}

type FriendListResponse struct {
	List  []FriendInfoResponse `json:"list"`
	Count int                  `json:"count"`
}

type FriendNoteRequest struct {
	Note     string `json:"note"`
	FriendID uint   `json:"friendId"`
	UserID   uint   `header:"User-ID"`
}

type FriendNoteResponse struct {
}

type ScarchInfo struct {
	UserID   uint   `json:"userId"`
	NickName string `json:"nickName"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"isFriend"` //是否是好友
}

type SearchRequest struct {
	Keyword string `form:"keyword"`
	UserID  uint   `header:"User-ID"`
	Online  bool   `form:"online,optional"`
	Page    int    `form:"page,optional"`
	Limit   int    `form:"limit,optional"`
}

type SearchResponse struct {
	List  []ScarchInfo `json:"list"`
	Count int64        `json:"count"`
}

type UserInfoRequest struct {
	UserID uint `header:"User-ID"`
	Role   int8 `header:"Role"`
}

type UserInfoResponse struct {
	UserID         uint            `json:"userId"`
	NickName       string          `json:"nickName"`
	Abstract       string          `json:"abstract"`
	Avatar         string          `json:"avatar"`
	RecallMsg      *string         `json:"recallMsg"`      //撤回消息的提示内容
	FriendOnline   bool            `json:"friendOnline"`   //好友上线提醒
	Sound          bool            `json:"sound"`          //声音提醒
	SecureLink     bool            `json:"secureLink"`     //安全链接
	SavePwd        bool            `json:"savePwd"`        //保存密码
	SearchUser     int8            `json:"searchUser"`     //别人搜索用户的方式
	Verify         int8            `json:"verify"`         //添加好友验证方式
	VerifyQuestion *VerifyQuestion `json:"verifyQuestion"` //好友验证问题
}

type UserInfoUpdateRequest struct {
	UserID         uint            `header:"User-ID"`
	NickName       *string         `json:"nickName,optional" user:"nick_name"`
	Abstract       *string         `json:"abstract,optional" user:"abstract"`
	Avatar         *string         `json:"avatar,optional" user:"avatar"`
	RecallMsg      *string         `json:"recallMsg,optional" user_conf:"recall_msg"`
	FriendOnline   *bool           `json:"friendOnline,optional" user_conf:"friend_online"`
	Sound          *bool           `json:"sound,optional" user_conf:"sound"`
	SecureLink     *bool           `json:"secureLink,optional" user_conf:"secure_link"`
	SavePwd        *bool           `json:"savePwd,optional" user_conf:"save_pwd"`
	SearchUser     *int8           `json:"searchUser,optional" user_conf:"search_user"`
	Verify         *int8           `json:"verify,optional" user_conf:"verify"`
	VerifyQuestion *VerifyQuestion `json:"verifyQuestion,optional" user_conf:"verify_question"`
}

type UserInfoUpdateResponse struct {
}

type VerifyQuestion struct {
	Question1 *string `json:"question1,optional" user_conf:"question1"`
	Answer1   *string `json:"answer1,optional" user_conf:"answer1"`
	Question2 *string `json:"question2,optional" user_conf:"question2"`
	Answer2   *string `json:"answer2,optional" user_conf:"answer2"`
	Question3 *string `json:"question3,optional" user_conf:"question3"`
	Answer3   *string `json:"answer3,optional" user_conf:"answer3"`
}
