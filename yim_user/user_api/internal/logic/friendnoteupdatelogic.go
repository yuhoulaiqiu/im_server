package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
