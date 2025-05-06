package handler

import (
	"common/model"
	"common/proto/admin"
	"common/utlis"
	"errors"
	"time"
)

func AdminLogin(in *administrators.AdminLoginReq) (*administrators.AdminLoginResp, error) {
	a := model.EbSystemAdmin{}
	acoount, err := a.GetAdminByAcoount(in.Account)
	if err != nil {
		return nil, errors.New("没有该账号")
	}
	if utlis.Md5(in.Password) != acoount.Pwd {
		return nil, errors.New("密码错误")
	}
	return &administrators.AdminLoginResp{
		AdminId: int64(acoount.Id),
	}, nil
}

func ProcessEnter(in *administrators.ProcessEnterReq) (*administrators.ProcessEnterResp, error) {
	ue := model.UserEnter{}
	//审核商家申请
	err := ue.UpdateStatus(in.MerchantId, in.Status)
	if err != nil {
		return nil, err
	}
	userEnter, err := ue.GetStatusById(in.MerchantId)
	if err != nil {
		return nil, err
	}

	ad := model.EbSystemAdmin{}
	admin, err := ad.GetAdminById(in.AdminId)
	if err != nil {
		return nil, err
	}
	//添加管理员操作记录表
	e := model.EbSystemLog{
		AdminId:    int(in.AdminId),
		AdminName:  admin.RealName,
		Path:       "admin/ump.storecombination/index/",
		Page:       "未知",
		Method:     "admin",
		Ip:         "127.0.0.1",
		Type:       "system",
		MerchantId: int(in.MerchantId),
	}
	err = e.Create()
	if err != nil {
		return nil, err
	}
	enter, err := ue.GetStatusById(in.MerchantId)
	if err != nil {
		return nil, err
	}
	switch enter.Status {
	case -1:
		return &administrators.ProcessEnterResp{Greet: false}, nil
	case 1:
		s := model.Stores{
			MerchantId:       in.MerchantId,
			StoreName:        userEnter.MerchantName,
			StoreDescription: userEnter.District,
			ContactPhone:     userEnter.LinkTel,
			Address:          userEnter.Address,
			CreatedAt:        time.Now(),
		}
		err = s.CreateStores()
		if err != nil {
			return nil, err
		}
		return &administrators.ProcessEnterResp{Greet: true}, nil
	default:
		return nil, nil
	}
}
