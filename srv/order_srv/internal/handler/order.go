package handler

import (
	"common/global"
	"common/model"
	"common/pkg"
	"common/proto/order"
	"errors"
	"github.com/google/uuid"
	"strconv"
)

func AddOrder(in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	pro := &model.Product{}
	err := pro.GetProductIdBy(in.ProductId)
	if err != nil {
		return nil, err
	}
	if pro.Id == 0 {
		return nil, errors.New("商品不存在")
	}
	if pro.IsShow == 0 {
		return nil, errors.New("商品下架")
	}
	if pro.Stock < uint32(in.Num) {
		return nil, errors.New("商品库存不足")
	}
	users := &model.User{}
	err = users.GetUserIdBy(in.Uid)
	if err != nil {
		return nil, err
	}
	if users.Status == 0 {
		return nil, errors.New("用户账号异常")
	}
	// 计算总金额
	totalPrice := float64(in.Num) * pro.Price
	// 判断优惠券是否存在
	cou := &model.Coupon{}
	err = cou.GetCouponIdBy(in.CouponId)
	if err != nil {
		return nil, err
	}
	if cou.Id == 0 {
		return nil, errors.New("该优惠券已下架")
	}
	// 计算实际金额
	var couponPrice float64
	var payPrice float64
	var deductionPrice float64
	if cou.CouponPrice <= totalPrice {
		couponPrice = cou.CouponPrice
		payPrice = totalPrice - cou.CouponPrice
	} else {
		couponPrice = 0
		payPrice = totalPrice
	}
	deductionPrice = totalPrice - payPrice
	// 计算积分
	gainIntegral := payPrice * 0.02
	// 开启事务
	tx := global.DB.Begin()
	err = pro.UpdateProductStock(in.ProductId, in.Num)
	if err != nil {
		return nil, errors.New("商品库存扣减失败")
	}
	orderSn := uuid.New().String() + strconv.Itoa(int(in.ProductId))
	orders := &model.Order{
		OrderSn:        orderSn,
		Uid:            uint32(in.Uid),
		RealName:       users.RealName,
		UserPhone:      users.Phone,
		UserAddress:    users.Addres,
		CartId:         uint32(in.CartId),
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
	}

	err = orders.AddOrder()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	price := strconv.FormatFloat(orders.PayPrice, 'f', 2, 64)
	payUrl := pkg.NewPay().Pay(pro.StoreName, orderSn, price)
	return &order.AddOrderResponse{PayUrl: payUrl}, nil
}

func PayCallback(in *order.PayCallbackRequest) (*order.PayCallbackResponse, error) {
	orders := model.Order{}
	status, _ := strconv.Atoi(in.Status)
	err := orders.UpdateOrderStatus(in.OrderSn, status)
	if err != nil {
		return nil, err
	}
	return &order.PayCallbackResponse{Success: true}, nil
}

func OrderList(in *order.OrderListRequest) (*order.OrderListResponse, error) {
	orders := &model.Order{}
	list, err := orders.OrderList(in.UserId, in.OrderStatus)
	if err != nil {
		return nil, err
	}
	var orderList []*order.OrderList
	for _, i := range list {
		op := &model.OrderProduct{}
		err = op.GetOrderProductIdBy(int64(i.Id))
		if err != nil {
			return nil, err
		}
		users := &model.User{}
		err = users.GetUserIdBy(int64(i.Uid))
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, &order.OrderList{
			OrderId:               int64(i.Id),
			OrderSn:               i.OrderSn,
			ProductId:             int64(op.ProductId),
			ProductName:           op.ProductName,
			ProductImage:          op.ProductImage,
			ProductSpecifications: op.ProductSpecifications,
			UserId:                int64(users.Uid),
			Account:               users.Account,
			UserPhone:             users.Phone,
			PayPrice:              float32(i.PayPrice),
			PayType:               int64(i.PayType),
			PayTime:               int64(i.PayTime),
			Paid:                  int64(i.Paid),
			Status:                int64(i.Status),
		})
	}
	return &order.OrderListResponse{List: orderList}, nil
}
