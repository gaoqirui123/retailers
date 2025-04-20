package router

import (
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup) {
	//o := r.Group("/order")
	{
		pkg.JWTAuth("retailers")
		//o.POST("/add", order.AddOrder)
		//o.POST("/callback", order.PayCallback)
	}
}
