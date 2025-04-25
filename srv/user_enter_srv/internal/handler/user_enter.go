package handler

import (
	"common/model"
	"common/pkg"
	"common/proto/user_enter"
	"errors"
	"regexp"
)

func Apply(in *user_enter.UserEnterApplyRequest) (*user_enter.UserEnterApplyResponse, error) {
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
	return &user_enter.UserEnterApplyResponse{Greet: "申请成功，等待平台审核"}, nil
}

func AddProduct(in *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
	p := model.Product{
		MerId:     in.MerId,
		Image:     in.Image,
		StoreName: in.StoreName,
		StoreInfo: in.StoreInfo,
		BarCode:   in.BarCode,
		CateId:    in.CateId,
		Price:     float64(in.Price),
		Postage:   float64(in.Postage),
		UnitName:  in.UnitName,
		Activity:  in.Activity,
	}
	err := p.Add()
	if err != nil {
		return nil, err
	}
	return &user_enter.AddProductResponse{Greet: "申请发布商品成功"}, nil
}

// AddCombinationProduct 商家发布拼团商品
func AddCombinationProduct(in *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
	p := model.Product{}
	product, err := p.GetProductById(in.ProductId, 0)
	if err != nil {
		return nil, err
	}
	c := &model.Combination{
		MerId:         int(in.MerId),
		ProductId:     int(in.ProductId),
		Image:         product.Image,
		Images:        product.Image,
		Title:         in.Title,
		Attr:          in.Attr,
		People:        int(in.People),
		Info:          product.StoreInfo,
		Price:         float64(in.Price),
		Sort:          int(in.Sort),
		Stock:         int(in.Stock),
		StartTime:     int(in.StartTime),
		StopTime:      int(in.StopTime),
		EffectiveTime: in.EffectiveTime,
		Cost:          int(product.Cost),
		UnitName:      product.UnitName,
		TempId:        in.TempId,
		Num:           in.Num,
		Quota:         in.Quota,
		QuotaShow:     in.QuotaShow,
	}
	err = c.Add()
	if err != nil {
		return nil, err
	}
	return &user_enter.AddCombinationProductResponse{Greet: "发布拼团商品成功"}, nil
}

// ProcessInvoice 商家审核用户的发票invoice申请
func ProcessInvoice(in *user_enter.ProcessInvoiceRequest) (*user_enter.ProcessInvoiceResponse, error) {
	/*
		申请审核：商家收到用户的发票申请后，对申请信息进行审核。审核内容包括用户消费记录、申请信息与系统内存储信息的一致性等。若信息无误，审核通过；若存在问题，会联系用户补充或修改信息。
	*/
	//1.判断用户的订单是否正常，是否付款,通过用户id和订单号查找订单

	//2.判断信息是否正确，价格，地址

	//3.同意申请/不同意申请，给出理由

	return &user_enter.ProcessInvoiceResponse{Greet: "审核成功"}, nil
}

func DelProduct(in *user_enter.DelProductRequest) (*user_enter.DelProductResponse, error) {
	p := model.Product{}
	product, err := p.GetProductById(in.MerId, in.Pid)
	if err != nil {
		return nil, err
	}
	if product.Id == 0 {
		return nil, errors.New("没有该商品")
	}
	//判断商品是否上架
	if product.IsShow == 0 {
		return nil, errors.New("该商品未上架")
	}
	err = product.UpdateStatus(in.Status, in.Pid)
	if err != nil {
		return nil, err
	}
	return &user_enter.DelProductResponse{Greet: true}, nil
}

func Register(in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	m := model.Merchant{
		MerchantAccount:  in.Account,
		MerchantPassword: in.Password,
		ContactPhone:     in.Phone,
		Email:            in.Email,
	}
	merchant, err := m.GetMerchantByAccount(in.Account)
	if merchant.MerchantId != 0 {
		return nil, errors.New("该账号已注册")
	}
	regex := `^1[3-9]\d{9}$`
	matchString, err := regexp.MatchString(regex, in.Phone)
	if !matchString {
		return nil, errors.New("手机号码不规范")
	}
	merchant, err = m.GetMerchantByPhone(in.Phone)
	if merchant.MerchantId != 0 {
		return nil, errors.New("该手机号已注册")
	}
	merchant, err = m.GetMerchantByEmail(in.Email)
	if merchant.MerchantId != 0 {
		return nil, errors.New("该邮箱已注册")
	}
	err = m.Create()
	if err != nil {
		return nil, errors.New("注册失败")
	}
	return &user_enter.UserEnterRegisterResponse{Greet: "注册成功"}, nil
}

func Login(in *user_enter.UserEnterLoginRequest) (*user_enter.UserEnterLoginResponse, error) {
	m := model.Merchant{
		MerchantAccount:  in.Account,
		MerchantPassword: in.Password,
	}
	merchant, _ := m.GetMerchantByAccount(in.Account)
	if merchant.MerchantId == 0 {
		return nil, errors.New("该账号未注册")
	}
	if in.Password != merchant.MerchantPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	claims := pkg.CustomClaims{
		ID: uint(merchant.MerchantId),
	}
	token, err := pkg.NewJWT("merchant").CreateToken(claims)
	if err != nil {
		return nil, err
	}
	return &user_enter.UserEnterLoginResponse{Greet: token}, nil
}
