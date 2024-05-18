package logic

import (
	"context"
	"yim_server/yim_user/user_models"

	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {
	var friends []user_models.FriendModel
	var count int64

	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit
	// 分页查询
	l.svcCtx.DB.Preload("RecvUserModel").Preload("SendUserModel").Limit(req.Limit).Offset(offset).Find(&friends, "send_user_id = ? or recv_user_id = ?", req.UserID, req.UserID)
	// 获取总记录数
	l.svcCtx.DB.Model(&friends).Where("send_user_id = ? or recv_user_id = ?", req.UserID, req.UserID).Count(&count)
	var list []types.FriendInfoResponse
	for _, friend := range friends {
		info := types.FriendInfoResponse{}
		if friend.SendUserId == req.UserID {
			// 发送者是自己
			info = types.FriendInfoResponse{
				UserID:   friend.RecvUserId,
				NickName: friend.RecvUserModel.NickName,
				Avatar:   friend.RecvUserModel.Avatar,
				Abstract: friend.RecvUserModel.Abstract,
				Note:     friend.SendNote,
			}
		} else if friend.RecvUserId == req.UserID {
			// 接收者是自己
			info = types.FriendInfoResponse{
				UserID:   friend.SendUserId,
				NickName: friend.SendUserModel.NickName,
				Avatar:   friend.SendUserModel.Avatar,
				Abstract: friend.SendUserModel.Abstract,
				Note:     friend.RecvNote,
			}
		}
		list = append(list, info)
	}
	return &types.FriendListResponse{List: list, Count: int(count)}, nil
}
