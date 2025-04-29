package handler

import (
	"common/global"
	"common/model"
	"common/proto/distribution"
	"common/utlis"
	"errors"
	"fmt"
	"strconv"
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
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("邀请吗无效")}, nil
	}

	//判断邀请码状态
	if id.Status == 2 {

		return &distribution.UserFillsInInvitationCodeResponse{Success: "邀请码已失效"}, nil

	}
	//对比邀请码
	if id.Code != in.Str {
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("邀请码有误")}, nil
	}

	u := model.User{
		Level: 1,
	}
	//查看我是否注册

	forMe, err := u.FindId(int(in.UserId))

	if err != nil {
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("%v", err)}, nil
	}

	//开启事务
	ctx := global.DB.Begin()

	//删除邀请码
	ctx.Begin()
	//未注册
	if forMe.Uid == 0 {

		u.Uid = int64(in.UserId)

		u.SpreadUid = id.Uid

		fmt.Println("确认上级用户id", id.Uid)

		u.Level = 2

		err = u.UserRegister()

		if err != nil {
			ctx.Rollback()
			return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("%v", err)}, nil
		}
	} else {

		//登录，确认上级id

		if !u.UpdatedSpreadUid(int(in.UserId), strconv.FormatInt(id.Uid, 10)) {
			ctx.Rollback()

			return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("%v", err)}, nil
		}

	}

	if !i.DeleteCode(in.Str) {

		ctx.Rollback()
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("邀请码有误")}, nil
	}

	//更改邀请码状态

	if !i.UpdateCode(in.Str) {

		ctx.Rollback()
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("更改状态有误")}, nil

	}

	err = ctx.Commit().Error
	if err != nil {
		return &distribution.UserFillsInInvitationCodeResponse{Success: fmt.Sprintf("事务提交有误%v", err)}, nil
	}

	return &distribution.UserFillsInInvitationCodeResponse{Success: "邀请码填写结束"}, nil
}

func DistributionLevelSetting(in *distribution.DistributionLevelSettingRequest) (*distribution.DistributionLevelSettingResponse, error) {
	dl := model.DistributionLevel{
		Img:       in.Img,
		LevelName: in.LevelName,
		Level:     in.Level,
		One:       float64(in.One),
		Two:       float64(in.Two),
	}
	//一级不能大于二级
	if in.One < in.Two {
		return &distribution.DistributionLevelSettingResponse{Success: false}, nil
	}

	if !dl.CreateDistributionLevel() {
		return &distribution.DistributionLevelSettingResponse{Success: false}, nil
	}

	return &distribution.DistributionLevelSettingResponse{Success: true}, nil
}

func TheCharts(in *distribution.TheChartsRequest) (*distribution.TheChartsResponse, error) {
	n := model.Commission{}
	list := n.CalculateAndRankTotalCommission()

	var sli []*distribution.ListRank
	for _, c := range list {
		sli = append(sli, &distribution.ListRank{
			ToUserId: int64(c.ToUserId),
			Amount:   float32(c.TotalAmount),
		})
	}
	return &distribution.TheChartsResponse{List: sli}, nil
}
