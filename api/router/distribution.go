package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func Distribution(c *gin.RouterGroup) {
	dis := c.Group("/distribution")
	{
		dis.Use(pkg.JWTAuth("retailers"))
		dis.POST("/add", handler.GenerateInvitationCode)
	}
}
