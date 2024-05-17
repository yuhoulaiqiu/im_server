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

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: int32(req.UserID),
	})
	if err != nil {
		return nil, err
	}
	var user user_models.UserModel
	err = json.Unmarshal(res.Data, &user)
	if err != nil {
		logx.Errorf("数据解析失败:%v", err)
		return nil, errors.New("数据解析失败")
	}
	resp = &types.UserInfoResponse{
		UserID:         user.ID,
		NickName:       user.NickName,
		Avatar:         user.Avatar,
		Abstract:       user.Abstract,
		RecallMsg:      user.UserConfModel.RecallMsg,
		FriendOnline:   user.UserConfModel.FriendOnline,
		Sound:          user.UserConfModel.Sound,
		SecureLink:     user.UserConfModel.SecureLink,
		SavePwd:        user.UserConfModel.SavePwd,
		SearchUser:     user.UserConfModel.SearchUser,
		Verify:         user.UserConfModel.Verify,
		VerifyQuestion: (*types.VerifyQuestion)(user.UserConfModel.VerifyQuestion),
	}
	if user.UserConfModel.VerifyQuestion != nil {
		resp.VerifyQuestion = &types.VerifyQuestion{
			Question1: user.UserConfModel.VerifyQuestion.Question1,
			Answer1:   user.UserConfModel.VerifyQuestion.Answer1,
			Question2: user.UserConfModel.VerifyQuestion.Question2,
			Answer2:   user.UserConfModel.VerifyQuestion.Answer2,
			Question3: user.UserConfModel.VerifyQuestion.Question3,
			Answer3:   user.UserConfModel.VerifyQuestion.Answer3,
		}
	}
	return resp, nil

}
