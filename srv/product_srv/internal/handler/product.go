package handler

import (
	"common/global"
	"common/model"
	"common/proto/product"
	"context"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

// CombinationList TODO:拼团商品列表展示
func CombinationList(in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
	c := model.Combination{}
	list, err := c.GetCombinationList()
	if err != nil {
		return nil, err
	}
	var lists []*product.CombinationList
	for _, combination := range list {
		l := product.CombinationList{
			Image:  combination.Images,
			Title:  combination.Title,
			People: int64(combination.People),
			Price:  float32(combination.Price),
			Stock:  int64(combination.Stock),
		}
		lists = append(lists, &l)
	}
	return &product.CombinationListResponse{List: lists}, nil
}

// GroupBuying TODO:用户发起拼团
func GroupBuying(in *product.GroupBuyingRequest) (*product.GroupBuyingResponse, error) {
	// 假设拼团时长为 1 小时，计算结束时间
	addtime := time.Now().Format("2006-01-02 15:04:05")
	stopTime := time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")
	c := model.Combination{}
	combination, err := c.GetCombinationById(in.Pid)
	if err != nil {
		return nil, err
	}
	orderId := uuid.New().String()
	u := model.User{}
	user, err := u.FindId(int(in.Uid))
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	totalPrice := float64(in.Num) * combination.Price
	/*	o := model.Order{
		OrderSn:        orderId,
		Uid:            in.Uid,
		RealName:       user.RealName,
		UserPhone:      user.Phone,
		UserAddress:    user.Address,
		FreightPrice:   float64(in.FreightPrice),
		TotalNum:       uint32(in.Num),
		TotalPrice:     totalPrice,
		PayPrice:       payPrice,
		DeductionPrice: deductionPrice,
		CouponId:       uint32(in.CouponId),
		CouponPrice:    couponPrice,
		Paid:           0,
		PayType:        uint32(in.PayType),
		GainIntegral:   gainIntegral,
		UseIntegral:    float64(in.UseIntegral),
		Mark:           in.Mark,
		MerId:          uint32(in.MerId),
		CombinationId:  uint32(in.CombinationId),
		PinkId:         uint32(in.PinkId),
		SeckillId:      uint32(in.ProductId),
		BargainId:      uint32(in.BargainId),
		StoreId:        int32(in.StoreId),
		ShippingType:   int8(in.ShippingType),
		IsChannel:      uint8(in.IsChannel),
	}*/
	p := model.Pink{
		Uid:        int(in.Uid),
		OrderId:    orderId,
		OrderIdKey: orderId,
		TotalNum:   int(in.Num),
		TotalPrice: totalPrice,
		Cid:        int(in.Pid),
		Pid:        combination.ProductId,
		People:     combination.People,
		Price:      combination.Price,
		AddTime:    addtime,
		StopTime:   stopTime,
	}
	err = p.Create()
	if err != nil {
		return nil, err
	}
	//将拼团信息存储到redis
	itoa := strconv.Itoa(p.Pid)
	key := "group_buy:" + itoa
	global.Rdb.Set(context.Background(), key, p, time.Hour)
	return &product.GroupBuyingResponse{Success: "发起拼团成功"}, nil
}
