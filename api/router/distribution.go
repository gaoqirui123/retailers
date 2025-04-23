package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func Distribution(c *gin.RouterGroup) {
	c.Use(pkg.JWTAuth("retailers"))
	d := c.Group("/distribution")
	{
		d.POST("/generateInvitationCode", handler.GenerateInvitationCode)
	}
}
