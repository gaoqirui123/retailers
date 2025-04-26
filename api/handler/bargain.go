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
	update, err := client.ProductUpdate(c, &product.ProductUpdateRequest{
		Id:        f.Id,
		IsBargain: f.IsBargain,
		UserID:    uint32(userId),
	})

	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", update)
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

}

// 砍价商品表列表
func BargainList(c *gin.Context) {

}

// 修改砍价商品表是否删除
func BargainUpdate(c *gin.Context) {

}
