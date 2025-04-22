package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/article"
	"github.com/gin-gonic/gin"
)

func ArticleRelease(c *gin.Context) {
	var data request.ArticleRelease
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}

	release, err := client.ArticleAdd(c, &article.ArticleAddRequest{

		Content:    data.Content,
		Cid:        data.Cid,
		Title:      data.Title,
		Author:     data.Author,
		ImageInput: data.ImageInput,
		Synopsis:   data.Synopsis,
		Hide:       data.Hide,
	})

	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", release)
}

func CategoryAdd(c *gin.Context) {
	var data request.CategoryAdd
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	release, err := client.CategoryAdd(c, &article.CategoryAddRequest{
		Pid:        data.Pid,
		Title:      data.Title,
		Intr:       data.Intr,
		ImageInput: data.ImageInput,
		Sort:       data.Sort,
		Status:     data.Status,
	})

	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", release)
}

func ArticleList(c *gin.Context) {
	release, err := client.ArticleList(c, &article.ArticleListRequest{})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", release)
}
