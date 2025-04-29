package handler

import (
	"common/global"
	"common/model"
	"common/pkg"
	"common/proto/order"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func AddOrder(in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	// 判断商品是否存在
	pro := &model.Product{}
	err := pro.GetProductIdBy(in.ProductId)
	if err != nil {
		return nil, err
	}
	if pro.Id == 0 {
		return nil, errors.New("商品不存在")
	}
	// 判断商品是否下架
	if pro.IsShow == 0 {
		return nil, errors.New("商品下架")
	}
	// 判断商品库存是否充足
	if pro.Stock < in.Num {
		return nil, errors.New("商品库存不足")
	}
	// 查询用户
	users := &model.User{}
	err = users.GetUserIdBy(in.Uid)
	if err != nil {
		return nil, err
	}
	if users.Status == 0 {
		return nil, errors.New("用户账号异常")
	}
	// 判断优惠券
	cou, err := JudgeCouponStatus(in.CouponId)
	if err != nil {
		return nil, err
	}
	// 计算总金额
	totalPrice := float64(in.Num) * pro.Price
	// 计算实际金额
	var payPrice float64
	var couponPrice float64
	if cou.CouponPrice <= totalPrice {
		couponPrice = cou.CouponPrice
		payPrice = totalPrice - cou.CouponPrice
	} else {
		couponPrice = 0
		payPrice = totalPrice
	}
	// 计算抵扣金额
	var deductionPrice float64
	deductionPrice = totalPrice - payPrice
	// 计算积分
	gainIntegral := payPrice * 0.02
	// 开启事务
	tx := global.DB.Begin()
	err = pro.UpdateProductStock(in.ProductId, in.Num)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("商品库存扣减失败")
	}
	orderSn := uuid.New().String() + strconv.Itoa(int(in.ProductId))
	orders := &model.Order{
		OrderSn:        orderSn,
		Uid:            in.Uid,
		RealName:       users.RealName,
		UserPhone:      users.Phone,
		UserAddress:    users.Address,
		FreightPrice:   pro.Postage,
		TotalNum:       in.Num,
		TotalPrice:     totalPrice,
		PayPrice:       payPrice,
		DeductionPrice: deductionPrice,
		CouponId:       in.CouponId,
		CouponPrice:    couponPrice,
		Paid:           0,
		PayType:        in.PayType,
		GainIntegral:   int64(gainIntegral),
		Mark:           in.Mark,
		MerId:          in.MerId,
		PinkId:         in.PinkId,
		SeckillId:      in.ProductId,
		BargainId:      in.BargainId,
		StoreId:        in.StoreId,
		ShippingType:   in.ShippingType,
		IsChannel:      in.IsChannel,
	}
	err = orders.AddOrder()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	op := &model.OrderProduct{
		OrderId:               orders.Id,
		ProductId:             in.ProductId,
		ProductName:           pro.StoreName,
		ProductImage:          pro.Image,
		ProductSpecifications: in.ProductSpecifications,
		ProductPrice:          pro.Price,
		ProductNum:            in.Num,
		ProductStatus:         pro.IsShow,
	}
	err = op.AddOrderProduct()
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

func JudgeCouponStatus(couponId int64) (*model.CouponUser, error) {
	// 判断优惠券是否存在
	cou := &model.CouponUser{}
	err := cou.GetCouponIdBy(couponId)
	if err != nil {
		return nil, err
	}
	if cou.Status == 2 {
		return nil, errors.New("该优惠券已过期")
	}
	if cou.Status == 1 {
		return nil, errors.New("该优惠券已使用")
	}
	return cou, nil
}

func PayCallback(in *order.PayCallbackRequest) (*order.PayCallbackResponse, error) {
	orders := &model.Order{}
	status, _ := strconv.Atoi(in.Status)
	if err := orders.UpdateOrderStatus(in.OrderSn, status); err != nil {
		return nil, err
	}
	timeData := time.Now().AddDate(0, 0, 0).Format("2006-01-02 15:04:05")
	err := orders.AddOrderPayTime(in.OrderSn, timeData)
	if err != nil {
		return nil, err
	}

	o := &model.Order{}
	od := o.GetOrderSnUserId(in.OrderSn)
	//查找不到消费用户
	if od.Id == 0 {
		return &order.PayCallbackResponse{Success: false}, err
	}
	//查找用户
	u := model.User{}
	id, err := u.FindId(int(od.Uid))
	if err != nil {
		return &order.PayCallbackResponse{Success: false}, err
	}
	var price float64
	dl := model.DistributionLevel{}
	fmt.Println(id)

	disLevel := dl.FindDistributionLevel(int(id.Level))

	fmt.Println("用户等级", disLevel.Level)

	if disLevel.Level == 1 {
		price = disLevel.One * float64(in.BuyerPayAmount)
	} else if disLevel.Level == 2 {
		price = disLevel.Two * float64(in.BuyerPayAmount)
	}
	n := &model.Commission{
		OrderSyn:   in.OrderSn,
		FromUserId: uint32(od.Uid),
		ToUserId:   uint32(id.SpreadUid),
		Level:      int8(id.Level),
		Amount:     price,
	}
	//同步返佣流水表
	if !n.CreateCommission() {
		return &order.PayCallbackResponse{Success: false}, nil
	}
	return &order.PayCallbackResponse{Success: true}, nil
}

func OrderList(in *order.OrderListRequest) (*order.OrderListResponse, error) {
	orders := &model.Order{}
	switch in.OrderStatus {
	case -1:
		list, err := orders.GetOrderStatusList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case -2:
		list, err := orders.GetOrderStatusList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 0:
		list, err := orders.AllOrderList(in.UserId)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 1:
		list, err := orders.GetOrderDelList(in.UserId, 1)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 2:
		list, err := orders.GetOrderPayList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 3:
		list, err := orders.GetOrderPayList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 4:
		list, err := orders.GetOrderStatusList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 5:
		list, err := orders.GetOrderStatusList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 6:
		list, err := orders.GetOrderStatusList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 7:
		list, err := orders.GetOrderStatusList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	case 8:
		list, err := orders.GetOrderStatusList(in.UserId, in.OrderStatus)
		if err != nil {
			return nil, err
		}
		orderList, err := OrderLists(list)
		if err != nil {
			return nil, err
		}
		return &order.OrderListResponse{List: orderList}, nil
	default:
		return nil, errors.New("无效状态")
	}
}

func OrderLists(list []*model.Order) ([]*order.OrderList, error) {
	var orderList []*order.OrderList
	for _, i := range list {
		op := &model.OrderProduct{}
		err := op.GetOrderProductIdBy(i.Id)
		if err != nil {
			return nil, err
		}
		users := &model.User{}
		err = users.GetUserIdBy(i.Uid)
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, &order.OrderList{
			OrderId:               i.Id,
			OrderSn:               i.OrderSn,
			ProductId:             op.ProductId,
			ProductName:           op.ProductName,
			ProductImage:          op.ProductImage,
			ProductSpecifications: op.ProductSpecifications,
			UserId:                users.Uid,
			Account:               users.Account,
			UserPhone:             users.Phone,
			PayPrice:              float32(i.PayPrice),
			PayType:               i.PayType,
			PayTime:               i.PayTime,
			Paid:                  i.Paid,
			Status:                i.Status,
		})
	}
	return orderList, nil
}
