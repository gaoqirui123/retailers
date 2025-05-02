package router

import (
	"api/handler"
	"common/pkg"

	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup) {
	o := r.Group("/order")

	{
		o.Use(pkg.JWTAuth("retailers"))
		o.POST("/add", handler.AddOrder)

		o.GET("/list", handler.OrderList)
		o.GET("/qrCodeVerification", handler.QrCodeVerification) //TODO:二维码核销
	}
}
