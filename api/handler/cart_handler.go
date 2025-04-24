package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/cart"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddCart(c *gin.Context) {
	var data request.AddCart
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	Uid := c.GetUint("userId")
	add, err := client.AddCart(c, &cart.AddCartRequest{
		Uid:               int64(Uid),
		Type:              data.Type,
		ProductId:         data.ProductId,
		ProductAttrUnique: data.ProductAttrUnique,
		CartNum:           data.CartNum,
		IsPay:             data.IsPay,
		IsNew:             data.IsNew,
		CombinationId:     data.CombinationId,
		SeckillId:         data.SeckillId,
		BargainId:         data.BargainId,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if add.CartId == 0 {
		response.RespError(c, "添加购物车失败")
		return
	}
	response.RespSuccess(c, "添加购物车成功", add.CartId)
}

func ClearCart(c *gin.Context) {
	Uid := c.GetUint("userId")
	clears, err := client.ClearCart(c, &cart.ClearCartRequest{
		Uid: uint64(Uid),
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if clears.Success == false {
		response.RespError(c, "清空购物车失败")
		return
	}
	response.RespSuccess(c, "清空购物车成功", clears.Success)
}

func DeleteCart(c *gin.Context) {
	var data request.DeleteCart
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	Uid := c.GetUint("userId")
	del, err := client.DeleteCart(c, &cart.DeleteCartRequest{
		Uid:       int64(Uid),
		ProductId: data.ProductId,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if del.Success == false {
		response.RespError(c, "删除购物车失败")
		return
	}
	response.RespSuccess(c, "删除购物车成功", del.Success)
}

func GetCartList(c *gin.Context) {
	Uid := c.GetUint("userId")
	list, err := client.GetCartList(c, &cart.GetCartListRequest{
		Uid: int64(Uid),
	})
	fmt.Println(list, "22222222222222")
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if list.List == nil {
		response.RespError(c, "购物车列表展示失败")
		return
	}
	response.RespSuccess(c, "购物车列表展示成功", list)
}
