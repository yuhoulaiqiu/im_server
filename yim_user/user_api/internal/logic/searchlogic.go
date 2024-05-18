package logic

import (
	"context"
	"yim_server/yim_user/user_models"

	"yim_server/yim_user/user_api/internal/svc"
	"yim_server/yim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	var users []user_models.UserModel
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit

	l.svcCtx.DB.Joins("UserConfModel").Where("UserConfModel.search_user != 0 AND (user_models.nick_name LIKE ? OR user_models.id = ?)", "%"+req.Keyword+"%", req.Keyword).Limit(req.Limit).Offset(offset).Find(&users)
	// 查询总数
	var count int64
	l.svcCtx.DB.Joins("UserConfModel").Where("UserConfModel.search_user != 0 AND (user_models.nick_name LIKE ? OR user_models.id = ?)", "%"+req.Keyword+"%", req.Keyword).Find(&users).Count(&count)
	list := make([]types.ScarchInfo, 0)
	for _, user := range users {
		var friend user_models.FriendModel
		isFriend := friend.IsFriend(l.svcCtx.DB, req.UserID, user.ID)
		list = append(list, types.ScarchInfo{
			UserID:   user.ID,
			NickName: user.NickName,
			Abstract: user.Abstract,
			Avatar:   user.Avatar,
			IsFriend: isFriend,
		})
	}
	return &types.SearchResponse{
		List:  list,
		Count: count,
	}, nil
}
