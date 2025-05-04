package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func UserEnter(r *gin.RouterGroup) {
	ue := r.Group("/userEnter")
	{
		ue.POST("/register", handler.Register) // TODO: 商户注册
		ue.POST("/login", handler.Login)       // TODO: 商户登录
		ue.Use(pkg.JWTAuth("merchant"))
		ue.POST("/apply", handler.Apply)                                // TODO: 商户申请
		ue.POST("/add/product", handler.AddProduct)                     // TODO: 商户发布商品
		ue.POST("/add/combination", handler.AddCombinationProduct)      // TODO: 发布拼团商品
		ue.POST("/update/status", handler.UpdateStatus)                 // TODO: 发布下架商品
		ue.POST("/invoice/process", handler.ProcessInvoice)             // TODO: 审核发票申请
		ue.GET("/invoice/list", handler.InvoiceList)                    // TODO: 发票列表展示
		ue.POST("/add/seckill", handler.AddSeckillProduct)              // TODO: 添加秒杀商品
		ue.POST("/reverse/stock", handler.ReverseStock)                 // TODO: 秒杀后反还剩余的商品
		ue.POST("/merchantVerification", handler.MerchantVerification)  // TODO: 商家核销
		ue.GET("/calculateOrderSummary", handler.CalculateOrderSummary) // TODO: 商家统计
		//ue.POST("/BatchReleaseOfProducts", handler.BatchReleaseOfProducts) // TODO: 商品批量发布
	}
}
