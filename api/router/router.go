package router

import (
	"api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/alipay", handler.PayCallback) //支付回调
	c := r.Group("/api/c")
	{
		UserRouter(c)
		ProductRouter(c)
		OrderRouter(c)
		ArticleRouter(c)
		CartRouter(c)
		Distribution(c)
	}
	b := r.Group("/api/b")
	{
		UserEnter(b)
		Administrators(b)
	}

}
