package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/user_enter"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var data request.Register
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	uid := c.GetUint("userId")
	register, err := client.Register(c, &user_enter.UserEnterRegisterRequest{
		Uid:          int64(uid),
		Province:     data.Province,
		City:         data.City,
		District:     data.District,
		Address:      data.Address,
		MerchantName: data.MerchantName,
		LinkTel:      data.LinkTel,
		Charter:      data.Charter,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "申请成功，等待平台审核", register)
}

func AddProduct(c *gin.Context) {
	var data request.AddProduct
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	uid := c.GetUint("userId")
	register, err := client.AddProduct(c, &user_enter.AddProductRequest{
		MerId:     int64(uid),
		Image:     data.Image,
		StoreName: data.StoreName,
		StoreInfo: data.StoreInfo,
		BarCode:   data.BarCode,
		CateId:    data.CateId,
		Price:     float32(data.Price),
		Postage:   float32(data.Postage),
		UnitName:  data.UnitName,
		Activity:  data.Activity,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "申请成功，等待平台审核", register)
}
