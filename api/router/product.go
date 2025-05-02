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
		p.POST("/group/buy", handler.GroupBuying)         //todo: 用户发起拼团
		p.POST("/group/join", handler.JoinGroupBuying)    //todo: 用户参与拼团
		p.POST("/add/seckill", handler.AddSeckillProduct) //TODO: 添加秒杀商品
		p.POST("/reverse/stock", handler.ReverseStock)    //TODO: 秒杀后反还剩余的商品
	}
}
