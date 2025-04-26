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

// 修改砍价商品表是否删除
func BargainUpdate(req *product.BargainUpdateRequest) (*product.BargainUpdateResponse, error) {
	return &product.BargainUpdateResponse{}, nil
}

// 砍价商品表详情
func BargainShow(req *product.BargainShowRequest) (*product.BargainShowResponse, error) {
	return &product.BargainShowResponse{}, nil
}

// 砍价商品表列表
func BargainList(req *product.BargainListRequest) (*product.BargainListResponse, error) {
	return &product.BargainListResponse{}, nil
}
