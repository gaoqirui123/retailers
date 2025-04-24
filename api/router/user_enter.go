package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func UserEnter(r *gin.RouterGroup) {
	ue := r.Group("/userEnter")
	{
		ue.Use(pkg.JWTAuth("retailers"))
		ue.POST("/register", handler.Register)                     //商户申请注册
		ue.POST("/add/product", handler.AddProduct)                //商户发布商品
		ue.POST("/add/combination", handler.AddCombinationProduct) //发布拼团商品
	}

}
