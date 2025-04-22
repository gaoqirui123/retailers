package router

import (
	"api/handler"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	p := r.Group("/product")
	{
		p.GET("/combination/list", handler.CombinationList)
	}
}
