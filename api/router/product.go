package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	p := r.Group("/product")
	{
		p.GET("/combination/list", handler.CombinationList) //todo:拼团商品列表展示
		p.Use(pkg.JWTAuth("retailers"))
		p.POST("/group/buy", handler.GroupBuying)      //todo: 用户发起拼团
		p.POST("/group/join", handler.JoinGroupBuying) //todo: 用户参与拼团
	}
}
