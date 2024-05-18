package logic

import (
	"context"
	"errors"
	"yim_server/yim_user/user_models"

	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendNoteUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendNoteUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendNoteUpdateLogic {
	return &FriendNoteUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendNoteUpdateLogic) FriendNoteUpdate(req *types.FriendNoteRequest) (resp *types.FriendNoteResponse, err error) {
	var friend user_models.FriendModel
	//判断是否是好友
	if !friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("不是好友关系")
	}
	l.svcCtx.DB.Find(&friend, "(send_user_id = ? and recv_user_id = ?) or (send_user_id = ? and recv_user_id = ?)", req.UserID, req.FriendID, req.FriendID, req.UserID)
	// 修改备注
	if req.UserID == friend.SendUserId {
		friend.SendNote = req.Note
	} else {
		friend.RecvNote = req.Note
	}
	l.svcCtx.DB.Save(&friend)
	return
}
