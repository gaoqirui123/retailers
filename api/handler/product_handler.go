package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/product"
	"github.com/gin-gonic/gin"
)

func CombinationList(c *gin.Context) {
	list, err := client.CombinationList(c, &product.CombinationListRequest{})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "拼团商品展示成功", list)
}

func GroupBuying(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.GroupBuy
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	list, err := client.GroupBuying(c, &product.GroupBuyingRequest{
		Uid: int64(userId),
		Pid: data.Pid,
		Num: data.Num,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "发起拼团成功", list)
}

func JoinGroupBuying(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.JoinGroupBuy
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	list, err := client.JoinGroupBuying(c, &product.JoinGroupBuyingRequest{
		Uid:    int64(userId),
		PinkId: data.PinkId,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "参与拼团成功", list)
}
