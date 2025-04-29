package handler

import (
	"retailers/api/client"
	"retailers/api/request"
	"retailers/api/response"
	"retailers/common/proto/user_enter"
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
	response.RespSuccess(c, "发布商品成功", register)
}
func AddCombinationProduct(c *gin.Context) {
	var data request.AddCombinationProduct
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	uid := c.GetUint("userId")
	product, err := client.AddCombinationProduct(c, &user_enter.AddCombinationProductRequest{
		MerId:         int64(uid),
		ProductId:     data.ProductId,
		Title:         data.Title,
		Attr:          data.Attr,
		People:        int32(data.People),
		Price:         float32(data.Price),
		Sort:          data.Sort,
		Stock:         data.Stock,
		StartTime:     data.StartTime,
		StopTime:      data.StopTime,
		EffectiveTime: data.EffectiveTime,
		TempId:        data.TempId,
		Num:           data.Num,
		Quota:         data.Quota,
		QuotaShow:     data.QuotaShow,
	})
	if err != nil {
		return
	}
	response.RespSuccess(c, "发布拼团商品成功", product)
}
func ProcessInvoice(c *gin.Context) {
	var data request.ProcessInvoice
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
}
