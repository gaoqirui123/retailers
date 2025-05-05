package handler

import (
	"common/global"
	"common/model"
	"common/proto/product"
	"fmt"
	"math/rand"
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
	err = m.GetProductShow(int64(req.Id))
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
		ProductId: req.ProductId,
	}
	err := b.BargainShow(req.ProductId)
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
		ProductId: req.ProductId,
		IsDel:     uint8(req.IsDel),
	}
	err := bargain.BargainUpdate()
	if err != nil {
		return nil, err
	}
	// 再次从数据库中查询确认更新后的数据
	err = bargain.BargainShow(req.ProductId)
	if err != nil {
		return nil, err
	}
	return &product.BargainUpdateResponse{
		Id:    bargain.Id,
		IsDel: uint32(bargain.IsDel),
	}, nil
}

// TODO:创建用户参与砍价接口
func BargainUserCreate(req *product.BargainUserCreateRequest) (*product.BargainUserCreateResponse, error) {
	var bargain model.Bargain
	err := bargain.BargainShowID(req.BargainId)
	if err != nil {
		return nil, err
	}
	// 检查库存是否充足
	if bargain.Stock <= 0 {
		return nil, fmt.Errorf("砍价商品库存不足")
	}
	var bargainUser model.BargainUser
	err = bargainUser.BargainUserShow(req.Uid, req.BargainId)
	if err != nil {
		return nil, err
	}
	if bargainUser.Id == 0 {
		// 如果用户没有参与记录，则创建新的参与记录
		bargainUser = model.BargainUser{
			Uid:             req.Uid,
			BargainId:       req.BargainId,
			BargainPriceMin: req.BargainPriceMin,
			FinalPrice:      bargain.Price, // 初始最终价格为原价
			AddTime:         uint32(time.Now().Unix()),
			EndTime:         uint32(time.Now().Add(time.Hour).Unix()), // 设置砍价截止时间为当前时间加1小时，可根据需求调整
		}
		err = bargainUser.BargainUserCreate()
		if err != nil {
			return nil, err
		}
	}
	// 检查是否已经砍到最低价
	if bargainUser.FinalPrice <= bargainUser.BargainPriceMin {
		return &product.BargainUserCreateResponse{
			Id:     bargainUser.Id,
			Status: uint32(bargainUser.Status),
		}, nil
	}
	// 模拟别人帮助砍价，随机生成砍价金额
	rand.Seed(time.Now().UnixNano())
	price := rand.Float64()*(bargain.BargainMaxPrice-bargain.BargainMinPrice) + bargain.BargainMinPrice
	// 确保最终价格不会低于最低价
	if bargainUser.FinalPrice-price < bargainUser.BargainPriceMin {
		price = bargainUser.FinalPrice - bargainUser.BargainPriceMin
	}
	bargainUser.Price = price
	bargainUser.FinalPrice -= price
	bargainUser.BargainPrice += price // 更新已砍总金额
	// 更新用户参与砍价记录
	err = global.DB.Save(&bargainUser).Error
	if err != nil {
		return nil, err
	}
	// 创建 BargainUserHelp 表记录
	bargainUserHelp := model.BargainUserHelp{
		Uid:           req.Uid, // 假设请求中包含帮助者的用户 ID
		BargainId:     req.BargainId,
		BargainUserId: bargainUser.Id,
		Price:         price,
		AddTime:       uint32(time.Now().Unix()),
		CurrentPrice:  bargainUser.FinalPrice,
	}
	err = global.DB.Create(&bargainUserHelp).Error
	if err != nil {
		return nil, err
	}
	// 检查是否在规定时间内砍到最低价
	if bargainUser.FinalPrice <= bargainUser.BargainPriceMin && time.Now().Unix() <= int64(bargainUser.EndTime) {
		// 砍价成功，修改砍价状态
		bargainUser.Status = model.BargainUserStatusSuccess
		err = global.DB.Save(&bargainUser).Error
		if err != nil {
			return nil, err
		}
	} else if time.Now().Unix() > int64(bargainUser.EndTime) {
		// 超过规定时间，砍价失败，修改砍价状态
		bargainUser.Status = model.BargainUserStatusFailed
		err = global.DB.Save(&bargainUser).Error
		if err != nil {
			return nil, err
		}
	}

	return &product.BargainUserCreateResponse{
		Id:     bargainUser.Id,
		Status: uint32(bargainUser.Status),
	}, nil
}

// TODO:用户参与砍价信息详情
func BargainUserShow(req *product.BargainUserShowRequest) (*product.BargainUserShowResponse, error) {
	var BargainUser model.BargainUser
	err := BargainUser.BargainUserShow(req.Uid, req.BargainId)
	if err != nil {
		return nil, err
	}
	return &product.BargainUserShowResponse{
		Id:              BargainUser.Id,
		Uid:             BargainUser.Uid,
		BargainId:       BargainUser.BargainId,
		BargainPriceMin: BargainUser.BargainPriceMin,
		BargainPrice:    BargainUser.BargainPrice,
		Price:           BargainUser.Price,
		FinalPrice:      BargainUser.FinalPrice,
		Status:          uint32(BargainUser.Status),
		AddTime:         BargainUser.AddTime,
		IsDel:           int32(BargainUser.IsDel),
	}, nil
}

// TODO:砍价帮助记录详情
func BargainUserHelpShow(req *product.BargainUserHelpShowRequest) (*product.BargainUserHelpShowResponse, error) {
	var BargainUserHelp model.BargainUserHelp
	err := BargainUserHelp.BargainUserHelpShow(req.Id)
	if err != nil {
		return nil, err
	}
	return &product.BargainUserHelpShowResponse{
		Id:            BargainUserHelp.Id,
		Uid:           BargainUserHelp.Uid,
		BargainId:     BargainUserHelp.BargainId,
		BargainUserId: BargainUserHelp.BargainUserId,
		Price:         BargainUserHelp.Price,
		AddTime:       BargainUserHelp.AddTime,
		IsSuccess:     uint32(BargainUserHelp.IsSuccess),
		CurrentPrice:  uint32(BargainUserHelp.CurrentPrice),
	}, nil
}

// TODO:用户参与砍价信息列表
func BargainUserList(req *product.BargainUserListRequest) (*product.BargainUserListResponse, error) {
	var BargainUser model.BargainUser
	list, err := BargainUser.BargainUserList()
	if err != nil {
		return nil, err
	}
	var List []*product.BargainUserList
	for _, v := range list {
		userList := product.BargainUserList{
			Id:              v.Id,
			Uid:             v.Uid,
			BargainId:       v.BargainId,
			BargainPriceMin: v.BargainPriceMin,
			BargainPrice:    v.BargainPrice,
			Price:           v.Price,
			FinalPrice:      v.FinalPrice,
			Status:          uint32(v.Status),
			AddTime:         v.AddTime,
			IsDel:           int32(v.IsDel),
		}
		List = append(List, &userList)
	}

	return &product.BargainUserListResponse{BargainUserList: List}, nil
}

// TODO:砍价帮助记录列表
func BargainUserHelpList(req *product.BargainUserHelpListRequest) (*product.BargainUserHelpListResponse, error) {
	var BargainUserHelp model.BargainUserHelp
	list, err := BargainUserHelp.BargainUserHelpList()
	if err != nil {
		return nil, err
	}
	var List []*product.BargainUserHelpList
	for _, v := range list {
		userList := product.BargainUserHelpList{
			Id:            v.Id,
			Uid:           v.Uid,
			BargainId:     v.BargainId,
			BargainUserId: v.BargainUserId,
			Price:         v.Price,
			AddTime:       v.AddTime,
			IsSuccess:     uint32(v.IsSuccess),
			CurrentPrice:  uint32(v.CurrentPrice),
		}
		List = append(List, &userList)
	}
	return &product.BargainUserHelpListResponse{BargainUserHelpList: List}, nil
}
