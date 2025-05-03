package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/product"
	"github.com/gin-gonic/gin"
)

// 修改商品砍价状态
func ProductUpdate(c *gin.Context) {
	var f request.ProductUpdate
	if err := c.ShouldBind(&f); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	userId := c.GetUint("userId")
	productUpdate, err := client.ProductUpdate(c, &product.ProductUpdateRequest{
		Id:        f.ProductId,
		IsBargain: f.IsBargain,
		UserID:    uint32(userId),
	})

	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", productUpdate)
}

// 创建砍价表
func BargainCreate(c *gin.Context) {
	var f request.BargainCreate
	if err := c.ShouldBind(&f); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	userId := c.GetUint("userId")
	create, err := client.BargainCreate(c, &product.BargainCreateRequest{
		UserID:          uint32(userId),
		ProductId:       f.ProductId,
		Title:           f.Title,
		Image:           f.Image,
		UnitName:        f.UnitName,
		Stock:           f.Stock,
		Images:          f.Images,
		Price:           f.Price,
		MinPrice:        f.MinPrice,
		Num:             f.Num,
		BargainMaxPrice: f.BargainMaxPrice,
		BargainMinPrice: f.BargainMinPrice,
		BargainNum:      f.BargainNum,
		Status:          f.Status,
		GiveIntegral:    f.GiveIntegral,
		Info:            f.Info,
		IsPostage:       f.IsPostage,
		Postage:         f.Postage,
		Rule:            f.Rule,
		StoreName:       f.StoreName,
		TempId:          f.TempId,
		Cost:            f.Cost,
	})

	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", create)
}

// 砍价商品表详情
func BargainShow(c *gin.Context) {
	var f request.BargainShow
	if err := c.ShouldBind(&f); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	show, err := client.BargainShow(c, &product.BargainShowRequest{ProductId: f.ProductId})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", show)
}

// 砍价商品表列表
func BargainList(c *gin.Context) {
	list, err := client.BargainList(c, &product.BargainListRequest{})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", list.BargainList)

}

// 修改砍价商品表是否删除
func BargainUpdate(c *gin.Context) {
	var f request.BargainUpdate
	if err := c.ShouldBind(&f); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	userId := c.GetUint("userId")
	update, err := client.BargainUpdate(c, &product.BargainUpdateRequest{
		ProductId: f.ProductId,
		IsDel:     f.IsDel,
		UserID:    uint32(userId),
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", update)
}

func BargainUserCreate(c *gin.Context) {
	var f request.BargainUserCreate
	if err := c.ShouldBind(&f); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	userId := c.GetUint("userId")
	create, err := client.BargainUserCreate(c, &product.BargainUserCreateRequest{
		Uid:             uint32(userId),
		BargainId:       f.BargainId,
		BargainPriceMin: f.BargainPriceMin,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", create)
}

func BargainUserShow(c *gin.Context) {
	var f request.BargainUserShow
	if err := c.ShouldBind(&f); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	userId := c.GetUint("userId")

	show, err := client.BargainUserShow(c, &product.BargainUserShowRequest{
		Uid:       uint32(userId),
		BargainId: f.BargainId,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", show)
}

func BargainUserHelpShow(c *gin.Context) {
	var f request.BargainUserHelpShow
	if err := c.ShouldBind(&f); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	show, err := client.BargainUserHelpShow(c, &product.BargainUserHelpShowRequest{Id: f.Id})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", show)
}

func BargainUserList(c *gin.Context) {
	list, err := client.BargainUserList(c, &product.BargainUserListRequest{})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", list.BargainUserList)
}

func BargainUserHelpList(c *gin.Context) {
	list, err := client.BargainUserHelpList(c, &product.BargainUserHelpListRequest{})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", list.BargainUserHelpList)
}
