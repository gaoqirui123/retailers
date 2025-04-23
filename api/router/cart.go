package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func CartRouter(c *gin.RouterGroup) {
	cart := c.Group("/cart")
	{
		cart.Use(pkg.JWTAuth("retailers"))
		cart.POST("/add", handler.AddCart)
		cart.POST("/clear", handler.ClearCart)
		cart.POST("/del", handler.DeleteCart)
		cart.GET("/list", handler.GetCartList)
	}
}
