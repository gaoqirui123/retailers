package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func UserEnter(r *gin.RouterGroup) {
	ue := r.Group("/userEnter")
	{
		ue.POST("/register", handler.Register) //商户注册
		ue.POST("/login", handler.Login)       //商户登录
		ue.Use(pkg.JWTAuth("merchant"))
		ue.POST("/apply", handler.Apply)                           //商户申请
		ue.POST("/add/product", handler.AddProduct)                //商户发布商品
		ue.POST("/add/combination", handler.AddCombinationProduct) //发布拼团商品
		ue.POST("/update/status", handler.UpdateStatus)            //发布下架商品
		ue.POST("/invoice/process", handler.ProcessInvoice)        //审核发票申请
		ue.GET("/invoice/list", handler.InvoiceList)               //发票列表展示
		//ue.POST("/BatchReleaseOfProducts", handler.BatchReleaseOfProducts) //发票列表展示
	}

}
