package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func Administrators(r *gin.RouterGroup) {
	ad := r.Group("/admin")
	{
		ad.POST("/login", handler.AdminLogin)
		ad.Use(pkg.JWTAuth("2209A"))
		ad.POST("/processEnter", handler.ProcessEnter)
	}

}
