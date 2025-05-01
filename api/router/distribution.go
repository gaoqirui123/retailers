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
		dis.POST("/add", handler.GenerateInvitationCode)                           //生成邀请码
		dis.POST("/user/fillsInInvitationCode", handler.UserFillsInInvitationCode) //用户填写邀请码
		dis.POST("/distribution_LevelSetting", handler.DistributionLevelSetting)   //分销等级设置
		dis.GET("/the_charts", handler.TheCharts)                                  //佣金排行榜
	}
}
