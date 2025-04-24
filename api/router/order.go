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
		o.POST("/callback", handler.PayCallback)
		o.GET("/list", handler.OrderList)
	}
}
