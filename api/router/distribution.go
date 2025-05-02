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
		dis.POST("/add", handler.GenerateInvitationCode)                           //TODO 生成邀请码
		dis.POST("/user/fillsInInvitationCode", handler.UserFillsInInvitationCode) //TODO 用户填写邀请码
		dis.POST("/distribution_LevelSetting", handler.DistributionLevelSetting)   // TODO 分销等级设置
		dis.GET("/the_charts", handler.TheCharts)                                  // TODO 佣金排行榜
		dis.GET("/lookDoneUp", handler.LookDoneUp)                                 //TODO 用户上下级展示

	}
}
