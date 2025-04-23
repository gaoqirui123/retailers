package handler

import (
	"common/model"
	"common/pkg"
	"common/proto/admin"
	"common/utlis"
	"errors"
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
	claims := pkg.CustomClaims{
		ID: uint(acoount.Id),
	}
	token, err := pkg.NewJWT("2209A").CreateToken(claims)
	if err != nil {
		return nil, err
	}
	return &administrators.AdminLoginResp{Greet: token}, nil
}

func ProcessEnter(in *administrators.ProcessEnterReq) (*administrators.ProcessEnterResp, error) {
	ue := model.UserEnter{}
	err := ue.UpdateStatus(in.MerchantId, in.Status)
	if err != nil {
		return nil, err
	}
	ad := model.EbSystemAdmin{}
	admin, err := ad.GetAdminById(in.AdminId)
	if err != nil {
		return nil, err
	}
	/*	atoi, err := strconv.Atoi(time.Now().AddDate(0, 0, 0).Format("20060102150405"))
		if err != nil {
			return nil, err
		}*/
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
		return &administrators.ProcessEnterResp{Greet: "审核未通过"}, nil
	case 1:
		return &administrators.ProcessEnterResp{Greet: "审核通过"}, nil
	default:
		return nil, nil
	}
}
