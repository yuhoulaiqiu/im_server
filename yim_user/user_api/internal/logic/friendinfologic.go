package logic

import (
	"context"
	"encoding/json"
	"errors"
	"yim_server/yim_user/user_models"
	"yim_server/yim_user/user_rpc/types/user_rpc"

	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	//是否是好友
	var friend user_models.FriendModel
	if !friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("ta不是你的好友哦~")
	}
	//获取好友信息
	res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: int32(req.FriendID),
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var user user_models.UserModel
	json.Unmarshal(res.Data, &user)
	resp = &types.FriendInfoResponse{
		UserID:   user.ID,
		NickName: user.NickName,
		Abstract: user.Abstract,
		Avatar:   user.Avatar,
	}
	resp.Note = friend.GetUserNote(req.UserID)

	return resp, nil
}
