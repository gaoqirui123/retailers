package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(c *gin.RouterGroup) {
	c.Use(pkg.JWTAuth("retailers"))
	a := c.Group("/article")
	{
		a.POST("/search", handler.ArticleAdd)                            //文章管理添加
		a.POST("/category/add", handler.CategoryAdd)                     //文章分类添加
		a.GET("/list", handler.ArticleList)                              //查询文章管理列表
		a.GET("/article/list", handler.CategoryList)                     //查询文章管理分类列表
		a.GET("/article/search", handler.ArticleSearch)                  //文章标题搜索
		a.POST("/article/edit", handler.EditArticle)                     //编辑文章
		a.GET("/delete/article", handler.DeleteArticle)                  //删除文章管理
		a.GET("/delete/article/category", handler.DeleteArticleCategory) //删除文章分类
	}
}
