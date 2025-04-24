package handler

import (
	"common/model"
	"common/proto/distribution"
	"common/utlis"
	"errors"
)

func GenerateInvitationCode(in *distribution.GenerateInvitationCodeRequest) (*distribution.GenerateInvitationCodeResponse, error) {
	// 获取用户头像
	u := model.User{}
	id, err := u.FindId(int(in.UserId))
	if err != nil {
		return nil, errors.New("用户查找失败")
	}
	if id.Uid == 0 {
		return nil, errors.New("用户查找失败")
	}

	var url string

	switch in.Type {

	case 1:
		str := utlis.GenerateInviteCode()
		url = str

	case 2:
		https := utlis.ChatUrl(in.UserId, id.Avatar)
		url = https
	default:
		url = "请选择邀请码方式"

	}
	return &distribution.GenerateInvitationCodeResponse{Url: url}, nil
}
