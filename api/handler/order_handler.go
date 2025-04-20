package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/order"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddOrder(c *gin.Context) {
	var data request.AddOrder
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	Uid := c.GetUint("userId")
	add, err := client.AddOrder(c, &order.AddOrderRequest{
		Uid:           int64(Uid),
		ProductId:     data.ProductId,
		Num:           data.Num,
		PayType:       data.PayType,
		CartId:        data.CartId,
		FreightPrice:  data.FreightPrice,
		CouponId:      data.CouponId,
		UseIntegral:   float32(data.UseIntegral),
		Mark:          data.Mark,
		StoreId:       data.StoreId,
		MerId:         data.MerId,
		BargainId:     data.BargainId,
		ShippingType:  data.ShippingType,
		IsChannel:     data.IsChannel,
		CombinationId: data.CombinationId,
		PinkId:        data.PinkId,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if add.PayUrl == "" {
		response.RespError(c, "创建订单失败")
		return
	}
	response.RespSuccess(c, "创建订单成功", add.PayUrl)
}

func PayCallback(c *gin.Context) {
	fmt.Println("支付回调")
	orderSn := c.Request.FormValue("out_trade_no")
	status := c.Request.FormValue("trade_status")
	callback, err := client.PayCallback(c, &order.PayCallbackRequest{
		OrderSn: orderSn,
		Status:  status,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if callback.Success == false {
		response.RespError(c, "支付回调失败")
		return
	}
	response.RespSuccess(c, "支付回调成功", callback)
}

func OrderList(c *gin.Context) {
	var data request.OrderList
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	UserId := c.GetUint("userId")
	list, err := client.OrderList(c, &order.OrderListRequest{
		UserId:      int64(UserId),
		OrderStatus: data.OrderStatus,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if list.List == nil {
		response.RespError(c, "订单列表展示失败")
		return
	}
	response.RespSuccess(c, "订单列表展示成功", list.List)
}
