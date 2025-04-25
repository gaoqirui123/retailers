package handler

import (
	"common/global"
	"common/model"
	"common/proto/distribution"
	"common/utlis"
	"errors"
	"fmt"
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

	i := model.InvitationCode{
		Uid:  in.UserId,
		Code: url,
	}
	if !i.Create() {
		return &distribution.GenerateInvitationCodeResponse{Url: fmt.Sprintf("%v", err)}, nil
	}

	return &distribution.GenerateInvitationCodeResponse{Url: fmt.Sprintf("邀请码生成%v", url)}, nil
}

func UserFillsInInvitationCode(in *distribution.UserFillsInInvitationCodeRequest) (*distribution.UserFillsInInvitationCodeResponse, error) {
	i := model.InvitationCode{}

	id := i.FindCode(in.Str)
	//查找邀请码
	if id.Id == 0 {
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("用户查找不存在")}, nil
	}
	//判断邀请码状态
	if id.Status == 2 {

		return &distribution.UserFillsInInvitationCodeResponse{Success: "邀请码已失效"}, nil

	}
	//对比邀请码
	if id.Code != in.Str {
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("邀请码有误")}, nil
	}
	//开启事务
	ctx := global.DB.Begin()
	//删除邀请码
	ctx.Begin()

	if !i.DeleteCode(in.Str) {
		ctx.Rollback()
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("邀请码有误")}, nil
	}

	//更改邀请码状态

	if !i.UpdateCode(in.Str) {
		ctx.Rollback()
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("邀请码有误")}, nil
	}

	return &distribution.UserFillsInInvitationCodeResponse{Success: "邀请码填写成功"}, nil
}
