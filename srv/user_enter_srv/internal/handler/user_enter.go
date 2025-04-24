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
	product, err := p.GetProductById(in.ProductId)
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
	//1.判断用户的订单是否正常，是否付款
	//2.判断信息是否正确，价格，地址
	//3.同意申请/不同意申请，给出理由
	return nil, nil
}
