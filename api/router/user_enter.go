package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func UserEnter(r *gin.RouterGroup) {
	ue := r.Group("/merchant")
	{
		ue.Use(pkg.JWTAuth("retailers"))
		ue.POST("./register", handler.Register)
		ue.POST("./add/product", handler.AddProduct)
	}

}
