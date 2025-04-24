package handler

import (
	"common/global"
	"common/model"
	"common/proto/product"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
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
	addtime := time.Now().Format(global.TimeFormat) //开始时间
	stopTime := time.Now().Add(time.Hour).Format(global.TimeFormat)
	//拼团商品表查询商品
	c := model.Combination{}
	combination, err := c.GetCombinationById(in.Pid)
	if err != nil {
		return nil, err
	}
	//生成订单id
	orderId := uuid.New().String()
	// 生成唯一的拼团 ID
	pinkId := rand.Intn(1000000)
	//用户表查找用户信息
	u := model.User{}
	user, err := u.FindId(int(in.Uid))
	if err != nil {
		return nil, err
	}
	//商品总价格
	totalPrice := float64(in.Num) * combination.Price
	o := model.Order{
		OrderSn:       orderId,
		Uid:           in.Uid,
		RealName:      user.RealName,
		UserPhone:     user.Phone,
		UserAddress:   user.Address,
		TotalNum:      in.Num,
		TotalPrice:    totalPrice,
		PayPrice:      totalPrice,
		PayType:       1,
		MerId:         int64(combination.MerId),
		CombinationId: in.Pid,
		PinkId:        int64(pinkId),
	}
	//创建订单
	err = o.AddOrder()
	if err != nil {
		return nil, err
	}
	p := model.Pink{
		Uid:        int(in.Uid),
		OrderId:    strconv.Itoa(pinkId),
		OrderIdKey: orderId,
		TotalNum:   int(in.Num),
		TotalPrice: totalPrice,
		Cid:        int(in.Pid),
		Pid:        combination.ProductId,
		People:     int64(combination.People),
		CurrentNum: 1,
		Price:      combination.Price,
		AddTime:    addtime,
		StopTime:   stopTime,
	}
	marshal, _ := json.Marshal(p)
	//添加拼团
	err = p.Create()
	if err != nil {
		return nil, err
	}
	//将拼团信息存储到redis
	key := global.GroupBuyKeyPrefix + strconv.Itoa(pinkId)
	err = global.Rdb.Set(context.Background(), key, marshal, time.Hour).Err()
	if err != nil {
		return nil, err
	}
	return &product.GroupBuyingResponse{Success: true}, nil
}

// JoinGroupBuying 用户参与拼团
func JoinGroupBuying(in *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
	ctx := context.Background()
	// 检查拼团是否存在
	key := "group_buy:" + in.PinkId
	exists, err := global.Rdb.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return nil, fmt.Errorf("拼团 %s 不存在", in.PinkId)
	}
	// 获取当前拼团信息
	groupInfoJSON, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("获取拼团信息失败: %w", err)
	}

	var pink model.Pink
	err = json.Unmarshal([]byte(groupInfoJSON), &pink)
	if err != nil {
		return nil, fmt.Errorf("反序列化拼团信息失败: %w", err)
	}
	// 检查拼团是否已结束
	endTime, err := time.Parse(global.TimeFormat, pink.StopTime)
	if err != nil {
		return nil, fmt.Errorf("解析拼团结束时间失败: %w", err)
	}
	if time.Now().After(endTime) {
		return nil, fmt.Errorf("拼团 %s 已结束，无法参与", in.PinkId)
	}
	// 检查拼团是否已满员
	if pink.CurrentNum >= pink.People {
		return nil, fmt.Errorf("拼团 %s 已完成，无法参与", in.PinkId)
	}

	// 更新拼团的当前人数
	pink.CurrentNum++
	pinkJSON, err := json.Marshal(pink)
	if err != nil {
		return nil, fmt.Errorf("序列化更新后的拼团信息失败: %w", err)
	}
	if err = global.Rdb.Set(ctx, key, pinkJSON, time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("更新拼团信息到 Redis 失败: %w", err)
	}

	// 检查拼团是否完成
	if pink.CurrentNum >= pink.People {

		err = pink.UpdateGroupStatus(key, 3)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("更新拼团状态失败: %w", err)
	}
	return &product.JoinGroupBuyingResponse{Success: true}, nil
}
