package handler

import (
	"common/global"
	"common/model"
	"common/pkg"
	"common/proto/order"
	"common/proto/product"
	"common/utlis"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
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
	//开启事务
	tx := global.DB.Begin()
	// 假设拼团时长为 1 小时，计算结束时间
	addtime := time.Now().Format(global.TimeFormat)                 //开始时间
	stopTime := time.Now().Add(time.Hour).Format(global.TimeFormat) //结束时间
	//拼团商品表查询商品
	c := model.Combination{}
	combination, err := c.GetCombinationById(in.Pid)
	if err != nil {
		return nil, err
	}
	// 检查库存
	if combination.Stock < int(in.Num) {
		tx.Rollback()
		return nil, fmt.Errorf("库存不足，当前库存: %d，需要数量: %d", combination.Stock, in.Num)
	}

	//生成订单id
	orderId := uuid.New().String()
	// 生成唯一的拼团 ID
	pinkId := rand.Intn(1000000)
	//用户表查找用户信息
	/*	u := model.User{}
		user, err := u.FindId(int(in.Uid))
		if err != nil {
			return nil, err
		}*/
	//商品总价格
	totalPrice := float64(in.Num) * combination.Price
	/*o := model.Order{
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
	}*/
	//创建订单

	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := order.NewOrderClient(conn)
	orderRequest := &order.AddOrderRequest{
		Uid:                   in.Uid,
		ProductId:             in.Pid,
		Num:                   in.Num,
		PayType:               1,
		CouponId:              0,
		Mark:                  "",
		StoreId:               0,
		MerId:                 int64(combination.MerId),
		BargainId:             1,
		ShippingType:          1,
		IsChannel:             0,
		PinkId:                int64(pinkId),
		ProductSpecifications: "白色",
	}
	addOrder, err := client.AddOrder(context.Background(), orderRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(addOrder)
	/*	err = o.AddOrder()
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	*/
	//扣mysql商品表总库存
	px := &model.Combination{}
	err = px.UpdateCombinationStock(in.Pid, in.Num)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("扣mysql商品表总库存失败")
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
	marshal, err := json.Marshal(p)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//添加拼团
	err = p.Create()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//将拼团信息存储到redis
	key := global.GroupBuyKeyPrefix + strconv.Itoa(pinkId)
	err = global.Rdb.Set(context.Background(), key, marshal, time.Hour).Err()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//发起支付
	pay := pkg.NewPay()
	sprintf := fmt.Sprintf("%.2f", totalPrice)
	s := pay.Pay(combination.Title, strconv.Itoa(pinkId), sprintf)
	fmt.Println(s)
	//生成拼团链接
	// 链接的基础部分
	baseURL := "https://314b3024.r39.cpolar.top/join_group"
	// 将拼团 ID 嵌入到链接中
	link := fmt.Sprintf("%s?id=%d", baseURL, pinkId)
	tx.Commit()
	return &product.GroupBuyingResponse{Success: link}, nil
}

// JoinGroupBuying TODO: 用户参与拼团
func JoinGroupBuying(in *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
	ctx := context.Background()
	// 检查拼团是否存在
	key := global.GroupBuyKeyPrefix + in.PinkId
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
	endTime, err := time.Parse(global.TimeFormat, pink.StopTime) //获取时间
	if err != nil {
		return nil, fmt.Errorf("解析拼团结束时间失败: %w", err)
	}
	if time.Now().After(endTime) {
		return nil, fmt.Errorf("拼团 %s 已结束，无法参与", in.PinkId)
	}
	// 检查拼团是否已满员
	if pink.CurrentNum == pink.People {
		return nil, fmt.Errorf("拼团 %s 已完成，无法参与", in.PinkId)
	}
	// 更新拼团的当前人数
	pink.CurrentNum++
	err = pink.UpdateGroupNum(in.PinkId, 1)
	if err != nil {
		return nil, err
	}
	pinkJSON, err := json.Marshal(pink)
	if err != nil {
		return nil, fmt.Errorf("序列化更新后的拼团信息失败: %w", err)
	}
	if err = global.Rdb.Set(ctx, key, pinkJSON, time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("更新拼团信息到 Redis 失败: %w", err)
	}
	// 更新拼团的状态，检查拼团是否完成1进行中2已完成3未完成
	if pink.CurrentNum == pink.People {
		err = pink.UpdateGroupStatus(context.Background(), key, 2)
		if err != nil {
			return nil, fmt.Errorf("更新拼团状态失败:%w", err)
		}
	}
	//生成订单id
	orderId := uuid.New().String()
	//用户表查找用户信息
	u := model.User{}
	user, err := u.FindId(int(in.Uid))
	if err != nil {
		return nil, err
	}
	//商品总价格
	totalPrice := pink.Price
	atoi, _ := strconv.Atoi(pink.OrderId)
	o := model.Order{
		OrderSn:       orderId,
		Uid:           in.Uid,
		RealName:      user.RealName,
		UserPhone:     user.Phone,
		UserAddress:   user.Address,
		TotalNum:      1,
		TotalPrice:    totalPrice,
		PayPrice:      totalPrice,
		PayType:       1,
		CombinationId: int64(pink.Pid),
		PinkId:        int64(atoi),
	}
	//创建订单
	err = o.AddOrder()
	if err != nil {
		return nil, err
	}
	//扣mysql商品表总库存
	px := &model.Combination{}
	err = px.UpdateCombinationStock(int64(pink.Pid), 1)
	if err != nil {
		return nil, errors.New("扣mysql商品表总库存失败")
	}
	//发起支付
	pay := pkg.NewPay()
	sprintf := fmt.Sprintf("%.2f", pink.Price)
	s := pay.Pay(pink.OrderIdKey, pink.OrderId, sprintf)
	return &product.JoinGroupBuyingResponse{Success: s}, nil
}

func AddSeckillProduct(in *product.AddSeckillProductRequest) (*product.AddSeckillProductResponse, error) {
	//查询商品是否存在
	p := &model.Product{}
	err := p.GetProductIdBy(in.ProductId)
	if err != nil {
		return nil, err
	}
	if p.Id == 0 {
		return nil, errors.New("商品不存在")
	}
	// 判断该商品是否是该商户的
	if p.MerId != in.UserEnterId {
		return nil, errors.New("该商品不是你的")
	}
	//判断商品库存不能小于秒杀库存
	if p.Stock < in.Num {
		return nil, errors.New("判断商品库存小于秒杀库存")
	}
	//开启事务
	tx := global.DB.Begin()
	//扣mysql商品表总库存
	err = p.UpdateProductStock(in.ProductId, in.Num)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("扣mysql商品表总库存失败")
	}
	seckill := &model.Seckill{
		MerId:       p.MerId,
		ProductId:   in.ProductId,
		Image:       p.Image,
		Images:      p.SliderImage,
		Name:        p.StoreName,
		Info:        p.StoreInfo,
		Price:       float64(in.Price),
		Cost:        p.Cost,
		OtPrice:     p.Price,
		Stock:       in.Num,
		Postage:     p.Postage,
		Description: in.Description,
		StartTime:   in.StartTime,
		StopTime:    in.StopTime,
		AddTime:     time.Now().Format(time.DateTime),
		Status:      p.IsShow,
		IsPostage:   p.IsPostage,
		Num:         in.Num,
		Quota:       in.Num,
		QuotaShow:   in.Num,
	}
	err = seckill.AddSeckillProduct()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if seckill.Id == 0 {
		tx.Rollback()
		return nil, errors.New("添加秒杀商品失败")
	}
	//将秒杀商品添加redis的list中
	utlis.ProductCreateRedis(int(seckill.Stock), int(seckill.Id))
	//判断redis库存是否添加成功
	val := utlis.GetProductRedis(int(seckill.Id))
	if val != seckill.Stock {
		tx.Rollback()
		return nil, errors.New("redis库存是否添加失败")
	}
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	return &product.AddSeckillProductResponse{SeckillId: seckill.Id}, nil
}

func ReverseStock(in *product.ReverseStockRequest) (*product.ReverseStockResponse, error) {
	// 查询秒杀商品是否存在
	s := &model.Seckill{}
	err := s.GetSeckillIdBY(in.SeckillId)
	if err != nil {
		return nil, err
	}
	if s.Id == 0 {
		return nil, errors.New("秒杀商品不存在")
	}
	// 判断该商品是否是该商户的
	if s.MerId != in.UserEnterId {
		return nil, errors.New("该商品不是你的")
	}
	//开启事务
	tx := global.DB.Begin()
	//反还剩余的商品
	g := &model.Product{}
	err = g.ReverseProductStock(s.ProductId, s.Stock)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("秒杀商品不存在")
	}
	//清除秒杀表里的数据
	err = s.DelSeckill(in.SeckillId)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("清除秒杀表里的数据失败")
	}
	//清除redis列表库存
	utlis.DelProductRedis(int(s.Id))
	//判断redis列表库存是否被清除
	val := utlis.GetProductRedis(int(s.Id))
	if val > 0 {
		tx.Rollback()
		return nil, errors.New("清除redis列表库存失败")
	}
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	return &product.ReverseStockResponse{Success: true}, nil
}
