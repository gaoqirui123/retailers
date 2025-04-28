package handler

import (
	"common/model"
	"common/proto/product"
	"fmt"
	"time"
)

// 修改商品表是否砍价状态
func ProductUpdate(req *product.ProductUpdateRequest) (*product.ProductUpdateResponse, error) {
	var m model.Product
	err := m.UpdateProductField(int(req.Id), req.IsBargain)
	if err != nil {
		return nil, err
	}
	//重新从数据库中查询更新后的记录
	err = m.GetProductIdBy(int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.ProductUpdateResponse{Id: uint32(m.Id), IsBargain: int32(m.IsBargain)}, nil
}

// 创建砍价商品信息
func BargainCreate(req *product.BargainCreateRequest) (*product.BargainCreateResponse, error) {
	currentTime := uint32(time.Now().Unix())
	bargain := model.Bargain{
		ProductId:       req.ProductId,
		Title:           req.Title,
		Image:           req.Image,
		UnitName:        req.UnitName,
		Stock:           int(req.Stock),
		Images:          req.Images,
		StartTime:       int(currentTime),
		StopTime:        int(currentTime + 3600),
		StoreName:       req.StoreName,
		Price:           req.Price,
		MinPrice:        req.MinPrice,
		Num:             int(req.Num),
		BargainMaxPrice: req.BargainMaxPrice,
		BargainMinPrice: req.BargainMinPrice,
		BargainNum:      int(req.BargainNum),
		Status:          uint8(req.Status),
		GiveIntegral:    req.GiveIntegral,
		Info:            req.Info,
		Cost:            req.Cost,
		AddTime:         int(currentTime),
		IsPostage:       uint8(req.IsPostage),
		Postage:         req.Postage,
		Rule:            req.Rule,
		TempId:          req.TempId,
	}
	// 创建砍价活动记录
	err := bargain.BargainCreate()
	if err != nil {
		return nil, fmt.Errorf("创建砍价活动记录失败: %v", err)
	}
	return &product.BargainCreateResponse{Id: bargain.Id}, nil
}

// 砍价商品表详情
func BargainShow(req *product.BargainShowRequest) (*product.BargainShowResponse, error) {
	b := model.Bargain{
		Id: req.Id,
	}
	err := b.BargainShow(req.Id)
	if err != nil {
		return nil, err
	}
	return &product.BargainShowResponse{
		Id:              b.Id,
		ProductId:       b.ProductId,
		Title:           b.Title,
		Image:           b.Image,
		UnitName:        b.UnitName,
		Stock:           int32(b.Stock),
		Sales:           int32(b.Sales),
		Images:          b.Images,
		StartTime:       int32(b.StartTime),
		StopTime:        int32(b.StopTime),
		StoreName:       b.StoreName,
		Price:           b.Price,
		MinPrice:        b.MinPrice,
		Num:             int32(b.Num),
		BargainMaxPrice: b.BargainMaxPrice,
		BargainMinPrice: b.BargainMinPrice,
		BargainNum:      int32(b.BargainNum),
		Status:          uint32(b.Status),
		GiveIntegral:    b.GiveIntegral,
		Info:            b.Info,
		Cost:            b.Cost,
		Sort:            int32(b.Sort),
		IsHot:           uint32(b.IsHot),
		IsDel:           uint32(b.IsDel),
		AddTime:         int32(b.AddTime),
		IsPostage:       uint32(b.IsPostage),
		Postage:         b.Postage,
		Rule:            b.Rule,
		Look:            int32(b.Look),
		Share:           int32(b.Share),
		TempId:          b.TempId,
		Weight:          b.Weight,
		Volume:          b.Volume,
		Quota:           b.Quota,
		QuotaShow:       b.QuotaShow,
	}, nil
}

// 砍价商品表列表
func BargainList(req *product.BargainListRequest) (*product.BargainListResponse, error) {
	var Bargain model.Bargain
	list, err := Bargain.BargainList()
	if err != nil {
		return nil, err
	}
	var List []*product.BargainList
	for _, v := range list {
		L := product.BargainList{
			Id:              v.Id,
			ProductId:       v.ProductId,
			Title:           v.Title,
			Image:           v.Image,
			UnitName:        v.UnitName,
			Stock:           int32(v.Stock),
			Sales:           int32(v.Sales),
			Images:          v.Images,
			StartTime:       int32(v.StartTime),
			StopTime:        int32(v.StopTime),
			StoreName:       v.StoreName,
			Price:           v.Price,
			MinPrice:        v.MinPrice,
			Num:             int32(v.Num),
			BargainMaxPrice: v.BargainMaxPrice,
			BargainMinPrice: v.BargainMinPrice,
			BargainNum:      int32(v.BargainNum),
			Status:          uint32(v.Status),
			GiveIntegral:    v.GiveIntegral,
			Info:            v.Info,
			Cost:            v.Cost,
			Sort:            int32(v.Sort),
			IsHot:           uint32(v.IsHot),
			IsDel:           uint32(v.IsDel),
			AddTime:         int32(v.AddTime),
			IsPostage:       uint32(v.IsPostage),
			Postage:         v.Postage,
			Rule:            v.Rule,
			Look:            int32(v.Look),
			Share:           int32(v.Share),
			TempId:          v.TempId,
			Weight:          v.Weight,
			Volume:          v.Volume,
			Quota:           v.Quota,
			QuotaShow:       v.QuotaShow,
		}
		List = append(List, &L)
	}

	return &product.BargainListResponse{BargainList: List}, nil
}

// 修改砍价商品表是否删除
func BargainUpdate(req *product.BargainUpdateRequest) (*product.BargainUpdateResponse, error) {
	bargain := model.Bargain{
		Id:    req.Id,
		IsDel: uint8(req.IsDel),
	}
	err := bargain.BargainUpdate()
	if err != nil {
		return nil, err
	}
	// 再次从数据库中查询确认更新后的数据
	err = bargain.BargainShow(req.Id)
	if err != nil {
		return nil, err
	}
	return &product.BargainUpdateResponse{
		Id:    bargain.Id,
		IsDel: uint32(bargain.IsDel),
	}, nil
}
