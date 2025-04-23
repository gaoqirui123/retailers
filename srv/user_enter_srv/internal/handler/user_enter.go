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
		MerId:     int(in.MerId),
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
