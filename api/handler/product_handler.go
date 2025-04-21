package handler

import (
	"api/client"
	"api/response"
	"common/proto/product"
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
