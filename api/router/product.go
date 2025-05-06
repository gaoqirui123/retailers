package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	p := r.Group("/product")
	{
		p.GET("/combination/list", handler.CombinationList)         //todo:拼团商品列表展示
		p.GET("/combination/info", handler.GetCombinationInfo)      //todo:拼团商品详情
		p.POST("/bargainShow", handler.BargainShow)                 //todo: 砍价商品表详情
		p.POST("/bargainList", handler.BargainList)                 //todo: 砍价商品表列表
		p.POST("/bargainUserHelpShow", handler.BargainUserHelpShow) //todo: 砍价帮助记录详情
		p.POST("/bargainUserList", handler.BargainUserList)         //todo: 用户参与砍价信息列表
		p.POST("/bargainUserHelpList", handler.BargainUserHelpList) //todo: 砍价帮助记录列表
		p.Use(pkg.JWTAuth("retailers"))
		p.POST("/group/buy", handler.GroupBuying)               //todo: 用户发起拼团
		p.POST("/group/join", handler.JoinGroupBuying)          //todo: 用户参与拼团
		p.POST("/bargainCreate", handler.BargainCreate)         //todo: 创建砍价表
		p.POST("/productUpdate", handler.ProductUpdate)         //todo: 修改商品砍价状态
		p.POST("/bargainUpdate", handler.BargainUpdate)         //todo: 修改砍价商品表是否删除
		p.POST("/bargainUserCreate", handler.BargainUserCreate) //todo: 创建用户参与砍价
		p.POST("/bargainUserShow", handler.BargainUserShow)     //todo: 用户参与砍价信息详情

	}
}
