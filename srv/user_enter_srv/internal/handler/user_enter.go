package handler

import (
	"common/model"
	"common/pkg"
	"common/proto/user_enter"
	"errors"
	"regexp"
	"time"
)

// Apply TODO:商家申请店铺
func Apply(in *user_enter.UserEnterApplyRequest) (*user_enter.UserEnterApplyResponse, error) {
	m := model.Merchant{}
	merchantById, err := m.GetMerchantById(in.UeId)
	if err != nil {
		return nil, err
	}
	ue := model.UserEnter{
		Uid:          int(in.UeId),
		Province:     in.Province,
		City:         in.City,
		District:     in.District,
		Address:      in.Address,
		MerchantName: in.MerchantName,
		LinkTel:      merchantById.ContactPhone,
		Charter:      in.Charter,
	}
	err = ue.Add()
	if err != nil {
		return nil, err
	}
	return &user_enter.UserEnterApplyResponse{Greet: true}, nil
}

// AddProduct TODO:商家发布商品
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
	return &user_enter.AddProductResponse{Greet: true}, nil
}

// AddCombinationProduct TODO:商家发布拼团商品
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
	return &user_enter.AddCombinationProductResponse{Greet: true}, nil
}

// ProcessInvoice TODO:商家审核用户的发票invoice申请
func ProcessInvoice(in *user_enter.ProcessInvoiceRequest) (*user_enter.ProcessInvoiceResponse, error) {
	// 检查输入参数是否为空
	if in == nil {
		return nil, errors.New("请求参数有误")
	}
	//查找用户是否有申请
	i := model.InvoiceApplication{}
	invoiceByUeId, err := i.GetInvoiceByUeId(in.Uid, in.OrderId)
	if err != nil {
		return nil, err
	}
	// 查找用户订单
	od := model.Order{}
	_, err = od.FindUserOrder(in.Uid, invoiceByUeId.OrderId)
	if err != nil {
		return nil, err
	}

	// 更新发票申请状态
	i = model.InvoiceApplication{}
	if in.Dis == "" {
		err = i.UpdateStatus(in.UeId, in.Uid, in.Status, in.OrderId, time.Now())
		if err != nil {
			return nil, errors.New("审核失败")
		}
	} else {
		err = i.UpdateStatusDis(in.UeId, in.Uid, in.Status, in.Dis, in.OrderId, time.Now())
		if err != nil {
			return nil, errors.New("审核失败")
		}
	}

	return &user_enter.ProcessInvoiceResponse{Greet: true}, nil
}

// InvoiceList TODO:发票列表展示
func InvoiceList(in *user_enter.InvoiceListRequest) (*user_enter.InvoiceListResponse, error) {
	ue := model.InvoiceApplication{}
	var lists []*user_enter.InvoiceList
	var applications []*model.InvoiceApplication
	var err error
	if in.Status == 0 {
		applications, err = ue.GetInvoiceByUeIds(in.UeId)
		if err != nil {
			return nil, errors.New("没有该商户的发票资料")
		}
	} else {
		applications, err = ue.GetInvoiceByUeIdAndStatus(in.UeId, in.Status)
		if err != nil {
			return nil, err
		}
	}

	for _, application := range applications {
		list := ue.ConvertToInvoiceList(application)
		lists = append(lists, &list)
	}

	return &user_enter.InvoiceListResponse{
		List: lists,
	}, nil
}

// DelProduct TODO:下架商品
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

// Register TODO:商家注册账号
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
	return &user_enter.UserEnterRegisterResponse{Greet: true}, nil
}

// Login TODO:商家账号登录
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
