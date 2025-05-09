package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
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
