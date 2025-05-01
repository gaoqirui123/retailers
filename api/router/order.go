package router

import (
	"api/handler"
	"common/pkg"

	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup) {
	o := r.Group("/order")
	o.POST("/alipay", handler.PayCallback) //支付回调
	{
		o.Use(pkg.JWTAuth("retailers"))
		o.POST("/add", handler.AddOrder)         // TODO: 创建订单
		o.POST("/callback", handler.PayCallback) // TODO: 支付回调
		o.GET("/list", handler.OrderList)        // TODO: 订单列表查询
	}
}
