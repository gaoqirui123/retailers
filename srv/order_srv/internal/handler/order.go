package handler

import (
	"common/cron"
	"common/global"
	"common/model"
	"common/pkg"
	"common/proto/order"
	"common/utlis"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type OrderProduct struct {
	Price       float64
	ProductName string
	Postage     float64
	Image       string
	IsShow      int64
}

func AddOrder(in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	// 检查用户状态
	users, err := checkUserStatus(in.Uid)
	if err != nil {
		return nil, err
	}

	// 检查商品状态
	pro, seckill, err := checkProductStatus(in)
	if err != nil {
		return nil, err
	}

	// 判断优惠券
	cou, err := JudgeCouponStatus(in.CouponId)
	if err != nil {
		return nil, err
	}

	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var orderProduct *OrderProduct
	if seckill.Id != 0 {
		// 扣减redis中的商品库存
		update := utlis.UpdateSeckillRedis(seckill.Id, in.Num)
		if !update {
			tx.Rollback()
			return nil, errors.New("redis商品库存扣减失败")
		}
		// 扣减mysql中的商品库存
		err = seckill.UpdateSeckillStock(seckill.Id, in.Num)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("商品库存扣减失败")
		}
		orderProduct = &OrderProduct{
			Price:       seckill.Price,
			ProductName: seckill.Name,
			Postage:     seckill.Postage,
			Image:       seckill.Image,
			IsShow:      seckill.IsShow,
		}
	}

	if pro.Id != 0 {
		// 扣减mysql中的商品库存
		err = pro.UpdateProductStock(seckill.Id, in.Num)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("商品库存扣减失败")
		}
		orderProduct = &OrderProduct{
			Price:       pro.Price,
			ProductName: pro.StoreName,
			Postage:     pro.Postage,
			Image:       pro.Image,
			IsShow:      pro.IsShow,
		}
	}

	// 计算总金额、实际金额和抵扣金额
	totalPrice, payPrice, deductionPrice, couponPrice := calculateOrderPrice(orderProduct.Price, cou.CouponPrice, in.Num)

	// 计算积分
	gainIntegral := payPrice * 0.02

	// 生成订单号
	orderSn := uuid.New().String() + strconv.Itoa(int(in.ProductId))
	addTime, _ := strconv.Atoi(time.Now().AddDate(0, 0, 0).Format("20060102"))
	orders := &model.Order{
		OrderSn:        orderSn,
		Uid:            in.Uid,
		RealName:       users.RealName,
		UserPhone:      users.Phone,
		UserAddress:    users.Address,
		FreightPrice:   orderProduct.Postage,
		TotalNum:       in.Num,
		TotalPrice:     totalPrice,
		PayPrice:       payPrice,
		DeductionPrice: deductionPrice,
		CouponId:       in.CouponId,
		CouponPrice:    couponPrice,
		Paid:           2,
		PayType:        in.PayType,
		AddTime:        int64(addTime),
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
		ProductName:           orderProduct.ProductName,
		ProductImage:          orderProduct.Image,
		ProductSpecifications: in.ProductSpecifications,
		ProductPrice:          orderProduct.Price,
		ProductNum:            in.Num,
		ProductStatus:         orderProduct.IsShow,
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

	// 起一个定时任务，查询30分钟后是否未支付，未支付的不扣库存
	go cron.OrderCron(orderSn)

	prices := strconv.FormatFloat(orders.PayPrice, 'f', 2, 64)
	fmt.Println(orderProduct.ProductName, orderSn, prices, ".............")
	payUrl := pkg.NewPay().Pay(orderProduct.ProductName, orderSn, prices)
	return &order.AddOrderResponse{PayUrl: payUrl}, nil
}

// 检查用户状态
func checkUserStatus(uid int64) (*model.User, error) {
	users := &model.User{}
	err := users.GetUserIdBy(uid)
	if err != nil {
		return nil, err
	}
	if users.Status == 0 {
		return nil, errors.New("账号异常无法下单")
	}
	return users, nil
}

// 检查商品状态
func checkProductStatus(in *order.AddOrderRequest) (*model.Product, *model.Seckill, error) {
	pro := &model.Product{}
	seckill := &model.Seckill{}
	switch in.Source {
	case 1:
		err := pro.GetProductIdBy(in.ProductId)
		if err != nil {
			return nil, nil, err
		}
		if pro.Id == 0 {
			return nil, nil, errors.New("商品不存在")
		}
		//if pro.IsShow == 0 {
		//	return nil, nil, errors.New("商品下架")
		//}
		if pro.Stock < in.Num {
			return nil, nil, errors.New("商品库存不足")
		}
	case 2:
		err := seckill.GetSeckillIdBY(in.ProductId)
		if err != nil {
			return nil, nil, err
		}
		if seckill.Id == 0 {
			return nil, nil, errors.New("秒杀商品不存在")
		}
		val := utlis.GetSeckillRedis(int(seckill.Id))
		if val != seckill.Stock {
			num := seckill.Stock - val
			utlis.SeckillCreateRedis(int(num), int(seckill.Id))
		}
		get := utlis.GetSeckillRedis(int(seckill.Id))
		if get != seckill.Stock {
			return nil, nil, errors.New("redis库存添加失败")
		}
	// 可以为其他 case 添加处理逻辑
	default:
		return nil, nil, errors.New("无效的商品来源")
	}
	return pro, seckill, nil
}

// 计算订单价格
func calculateOrderPrice(price, couponPrice float64, num int64) (float64, float64, float64, float64) {
	totalPrice := float64(num) * price
	var payPrice float64
	if couponPrice <= totalPrice {
		payPrice = totalPrice - couponPrice
	} else {
		couponPrice = 0
		payPrice = totalPrice
	}
	deductionPrice := totalPrice - payPrice
	return totalPrice, payPrice, deductionPrice, couponPrice
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
	var status int
	if in.Status == "WAIT_BUYER_PAY" {
		status = 1
	}
	if in.Status == "TRADE_CLOSED" {
		status = 2
	}
	if in.Status == "TRADE_SUCCESS" {
		status = 3
	}
	if in.Status == "TRADE_FINISHED" {
		status = 4
	}
	orders := &model.Order{}
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
	//查找用户等级
	u := model.User{}
	id, err := u.FindId(int(od.Uid))
	if err != nil {
		return &order.PayCallbackResponse{Success: false}, err
	}
	var price float64

	//查找配置的返利等级
	dl := &model.DistributionLevel{}

	disLevel := dl.FindDistributionLevel(int(id.Level))

	fmt.Println("用户等级", disLevel.Level)

	if disLevel.Level == 1 {
		price = disLevel.One * float64(in.BuyerPayAmount)
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
	} else if disLevel.Level == 2 {
		price = disLevel.Two * float64(in.BuyerPayAmount)
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

func QrCodeVerification(in *order.QrCodeVerificationRequest) (*order.QrCodeVerificationResponse, error) {
	o := &model.Order{}
	id, err := o.FindId(in.OrderId)
	if err != nil {
		return nil, err
	}
	////判断订单是否付款
	//
	//if id.Paid != 3 {
	//	return &order.QrCodeVerificationResponse{Success: false}, err
	//}
	////判断订单状态
	//
	//if id.Status != 5 {
	//	return &order.QrCodeVerificationResponse{Success: false}, err
	//}
	// 将订单信息序列化为 JSON 字符串
	Order := global.Order
	Order.Id = id.Id
	Order.OrderSn = id.OrderSn
	Order.Uid = id.Uid
	Order.Paid = id.Paid
	Order.Status = id.Status

	orderInfo, err := json.Marshal(Order)
	//	err = global.Rdb.Set(context.Background(), fmt.Sprintf(global.IMGName, in.UserId, in.OrderId), string(orderInfo), time.Minute*5).Err()
	//	if err != nil {
	//		return nil, err
	//	}

	if err != nil {
		return &order.QrCodeVerificationResponse{Success: err.Error()}, err
	}

	// 指定具体的文件名
	filePath := "../../common/img/" + fmt.Sprintf(global.IMGName, in.UserId, in.OrderId) + ".png"

	// 确保目录存在
	dir := "../../common/img"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	}
	logoPath := "../../common/img/1234.png" // 替换为你的 logo 图片路径
	err = utlis.GenerateQRCodeWithLogo(string(orderInfo), logoPath, filePath)
	if err != nil {
		return &order.QrCodeVerificationResponse{Success: err.Error()}, err
	}
	code, err := utlis.DecodeQRCode(filePath)
	if err != nil {
		return nil, err
	}

	all := json.Unmarshal([]byte(code), &Order)
	if all != nil {

		return &order.QrCodeVerificationResponse{Success: fmt.Sprintf("JSON 反序列化失败: %v\n", all)}, all
	}

	keyall := fmt.Sprintf(global.IMGName, Order.Uid, Order.Id)
	return &order.QrCodeVerificationResponse{Success: keyall}, nil
}
