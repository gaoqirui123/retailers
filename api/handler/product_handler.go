package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/product"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CombinationList(c *gin.Context) {
	list, err := client.CombinationList(c, &product.CombinationListRequest{})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "拼团商品展示成功", list)
}

func GroupBuying(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.GroupBuy
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	list, err := client.GroupBuying(c, &product.GroupBuyingRequest{
		Uid: int64(userId),
		Pid: data.Pid,
		Num: data.Num,
	})
	if err != nil {
		response.RespError(c, fmt.Sprintf(err.Error()))
		return
	}
	response.RespSuccess(c, "发起拼团成功", list)
}

func JoinGroupBuying(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.JoinGroupBuy
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	list, err := client.JoinGroupBuying(c, &product.JoinGroupBuyingRequest{
		Uid:    int64(userId),
		PinkId: data.PinkId,
	})
	if err != nil {
		response.RespError(c, fmt.Sprintf(err.Error()))
		return
	}
	response.RespSuccess(c, "参与拼团成功", list)
}

func AddSeckillProduct(c *gin.Context) {
	userEnterId := c.GetUint("userId")
	var data request.AddSeckillProduct
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	seckill, err := client.AddSeckillProduct(c, &product.AddSeckillProductRequest{
		UserEnterId: int64(userEnterId),
		ProductId:   data.ProductId,
		Num:         data.Num,
		Price:       float32(data.Price),
		Description: data.Description,
		StartTime:   data.StartTime,
		StopTime:    data.StopTime,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if seckill.SeckillId == 0 {
		response.RespError(c, "添加秒杀商品成功")
		return
	}
	response.RespSuccess(c, "添加秒杀商品成功", seckill)
}

func ReverseStock(c *gin.Context) {
	userEnterId := c.GetUint("userId")
	var data request.ReverseStock
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	reverse, err := client.ReverseStock(c, &product.ReverseStockRequest{
		UserEnterId: int64(userEnterId),
		SeckillId:   data.ProductId,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if reverse.Success == false {
		response.RespError(c, "秒杀后反还剩余的商品失败")
		return
	}
	response.RespSuccess(c, "秒杀后反还剩余的商品成功", reverse)
}
