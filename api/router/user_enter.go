package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func UserEnter(r *gin.RouterGroup) {
	ue := r.Group("/user")
	{
		ue.Use(pkg.JWTAuth("retailers"))
		ue.POST("./register", handler.Register)
	}

}
