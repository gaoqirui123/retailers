package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/article"
	"github.com/gin-gonic/gin"
)

func ArticleAdd(c *gin.Context) {
	var data request.ArticleAdd
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	userId := c.GetUint("userId")
	release, err := client.ArticleAdd(c, &article.ArticleAddRequest{
		Content:    data.Content,
		Cid:        data.Cid,
		Title:      data.Title,
		Author:     data.Author,
		ImageInput: data.ImageInput,
		Synopsis:   data.Synopsis,
		Hide:       data.Hide,
		Uid:        uint32(userId),
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

func CategoryAdd(c *gin.Context) {

	var data request.CategoryAdd
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	userId := c.GetUint("userId")
	release, err := client.CategoryAdd(c, &article.CategoryAddRequest{
		Pid:        data.Pid,
		Title:      data.Title,
		Intr:       data.Intr,
		ImageInput: data.ImageInput,
		Sort:       data.Sort,
		Status:     data.Status,
		Uid:        uint32(userId),
	})

	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

func ArticleList(c *gin.Context) {
	release, err := client.ArticleList(c, &article.ArticleListRequest{})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

func CategoryList(c *gin.Context) {
	var data request.CategoryList
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	release, err := client.CategoryList(c, &article.CategoryListRequest{
		Cid: data.Cid,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

func ArticleSearch(c *gin.Context) {
	var data request.ArticleSearch
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	release, err := client.ArticleSearch(c, &article.ArticleSearchRequest{
		Title: data.Title,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

func EditArticle(c *gin.Context) {
	var data request.EditArticle
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	release, err := client.EditArticle(c, &article.EditArticleRequest{
		Id:         data.Id,
		Cid:        data.Cid,
		Title:      data.Title,
		Author:     data.Author,
		ImageInput: data.ImageInput,
		Synopsis:   data.Synopsis,
		Hide:       data.Hide,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

func DeleteArticle(c *gin.Context) {
	var data request.DeleteArticle
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	release, err := client.DeleteArticle(c, &article.DeleteRequest{
		Ids: int64(data.Id),
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

func DeleteArticleCategory(c *gin.Context) {
	var data request.DeleteArticle
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	release, err := client.DeleteArticleCategory(c, &article.DeleteRequest{
		Ids: int64(data.Id),
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

// 发布评论
func PostAComment(c *gin.Context) {
	var data request.PostAComment
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	userId := c.GetUint("userId")
	release, err := client.PostAComment(c, &article.PostACommentRequest{
		Uid:         int64(userId),
		ArticleID:   int64(data.ArticleID),
		Content:     data.Content,
		Pid:         int64(data.Pid),
		ReplyUserID: int64(data.ReplyUserID),
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

// 文章点赞
func ArticleThumbsUp(c *gin.Context) {
	var data request.ArticleThumbsUp
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	userId := c.GetUint("userId")
	release, err := client.ArticleThumbsUp(c, &article.ArticleThumbsUpRequest{
		Uid:       int64(userId),
		ArticleID: int64(data.ArticleID),
		Button:    int64(data.Button),
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

// 删除评论
func DeleteComment(c *gin.Context) {
	var data request.DeleteComment
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	userId := c.GetUint("userId")
	release, err := client.DeleteComment(c, &article.DeleteCommentRequest{
		Uid:       int64(userId),
		ArticleID: int64(data.ArticleID),
		CommentID: int64(data.CommentID),
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}

// 高赞文章排序
func TopLikeArticleRanking(c *gin.Context) {
	var data request.TopLikeArticleRanking
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}

	release, err := client.TopLikeArticleRanking(c, &article.TopLikeArticleRankingRequest{
		Top: int64(data.Top),
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "成功", release)
}
