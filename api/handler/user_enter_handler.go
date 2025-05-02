package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/user_enter"

	"github.com/gin-gonic/gin"
)

func Apply(c *gin.Context) {
	var data request.Apply
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	uid := c.GetUint("userId")
	register, err := client.Apply(c, &user_enter.UserEnterApplyRequest{
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
func Register(c *gin.Context) {
	var data request.Register
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	register, err := client.Register(c, &user_enter.UserEnterRegisterRequest{
		Account:  data.Account,
		Password: data.Password,
		Phone:    data.Phone,
		Email:    data.Email,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "注册成功", register)
}
func Login(c *gin.Context) {
	var data request.Login
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	register, err := client.Login(c, &user_enter.UserEnterLoginRequest{
		Account:  data.Account,
		Password: data.Password,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "登录成功", register)
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

// ProcessInvoice TODO:审核发票申请
func ProcessInvoice(c *gin.Context) {
	var data request.ProcessInvoice
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	uid := c.GetUint("userId")
	invoice, err := client.ProcessInvoice(c, &user_enter.ProcessInvoiceRequest{
		UeId:   int64(uid),
		Uid:    data.Uid,
		Status: data.Status,
		Dis:    data.Dis,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "审核完成", invoice)
}

// UpdateStatus TODO: 下架商品
func UpdateStatus(c *gin.Context) {
	var data request.DelProduct
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	uid := c.GetUint("userId")
	product, err := client.DelProduct(c, &user_enter.DelProductRequest{
		MerId:  int64(uid),
		Pid:    data.Pid,
		Status: data.Status,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "下架商品成功", product)
}

// InvoiceList TODO:发票列表展示
func InvoiceList(c *gin.Context) {
	var data request.InvoiceList
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	uid := c.GetUint("userId")
	list, err := client.InvoiceList(c, &user_enter.InvoiceListRequest{
		UeId:   int64(uid),
		Status: data.Status,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "展示成功", list)
}

//批量添加商品

//func BatchReleaseOfProducts(c *gin.Context) {
//	var data request.BatchReleaseOfProducts
//	err := c.ShouldBind(&data)
//	if err != nil {
//		response.RespError(c, err.Error())
//		return
//	}
//	uid := c.GetUint("userId")
//	list, err := client.BatchReleaseOfProducts(c, &user_enter.BatchReleaseOfProductsRequest{
//		MerId:       int64(uid),
//		Image:       data.Image,
//		SliderImage: data.SliderImage,
//		StoreName:   data.StoreName,
//		CateId:      data.CateId,
//		IsShow:      data.IsShow,
//		Price:       int64(data.Price),
//		Postage:     int64(data.Postage),
//		UnitName:    data.UnitName,
//	})
//	if err != nil {
//		response.RespError(c, err.Error())
//		return
//	}
//	response.RespSuccess(c, "批量发布成功", list)
//}
