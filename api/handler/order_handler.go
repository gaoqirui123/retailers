package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/order"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddOrder(c *gin.Context) {
	var data request.AddOrder
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	Uid := c.GetUint("userId")
	add, err := client.AddOrder(c, &order.AddOrderRequest{
		Uid:          int64(Uid),
		ProductId:    data.ProductId,
		Num:          data.Num,
		PayType:      data.PayType,
		CouponId:     data.CouponId,
		Mark:         data.Mark,
		StoreId:      data.StoreId,
		MerId:        data.MerId,
		BargainId:    data.BargainId,
		ShippingType: data.ShippingType,
		IsChannel:    data.IsChannel,
		PinkId:       data.PinkId,
		Source:       data.Source,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	if add.PayUrl == "" {
		response.RespError(c, 500, "创建订单失败")
		return
	}
	response.RespSuccess(c, 200, "创建订单成功", add.PayUrl)
}

func PayCallback(c *gin.Context) {

	orderSn := c.Request.FormValue("out_trade_no")

	status := c.Request.FormValue("trade_status")

	buyerPayAmount := c.Request.FormValue("buyer_pay_amount") //用户在交易中支付的金额。

	fmt.Println("支付宝回调", status, orderSn, buyerPayAmount)
	price, err := strconv.ParseFloat(buyerPayAmount, 64)
	if err != nil {
		// 处理转换错误
		fmt.Println("Error converting buyer_pay_amount to float64:", err)
		return
	}

	callback, err := client.PayCallback(c, &order.PayCallbackRequest{
		BuyerPayAmount: float32(price),
		OrderSn:        orderSn,
		Status:         status,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	if callback.Success == false {
		response.RespError(c, 500, "支付回调失败")
		return
	}
	response.RespSuccess(c, 200, "支付回调成功", callback)
}

func OrderList(c *gin.Context) {
	var data request.OrderList
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	UserId := c.GetUint("userId")
	list, err := client.OrderList(c, &order.OrderListRequest{
		UserId:      int64(UserId),
		OrderStatus: data.OrderStatus,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	if list.List == nil {
		response.RespError(c, 500, "订单列表展示失败")
		return
	}
	response.RespSuccess(c, 200, "订单列表展示成功", list.List)
}

func QrCodeVerification(c *gin.Context) {
	var data request.QrCodeVerification
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	UserId := c.GetUint("userId")
	list, err := client.QrCodeVerification(c, &order.QrCodeVerificationRequest{
		UserId:  int64(UserId),
		OrderId: data.OrderId,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}

	response.RespSuccess(c, 200, "二维码已生成欢迎到点核销", list)

}
