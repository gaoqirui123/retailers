package handler

import (
	"common/model"
	"common/proto/user_enter"
	"errors"
)

func Register(in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	ue := model.UserEnter{
		Uid:          int(in.Uid),
		Province:     in.Province,
		City:         in.City,
		District:     in.District,
		Address:      in.Address,
		MerchantName: in.MerchantName,
		LinkTel:      in.LinkTel,
		Charter:      in.Charter,
	}
	if in.MerchantName == "" {
		return nil, errors.New("商户名称不能为空")
	}
	if in.LinkTel == "" {
		return nil, errors.New("商户电话不能为空")
	}
	err := ue.Add()
	if err != nil {
		return nil, err
	}
	return &user_enter.UserEnterRegisterResponse{Greet: "申请成功，等待平台审核"}, nil
}
