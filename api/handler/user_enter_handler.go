package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/pkg"
	"common/proto/user_enter"
	"github.com/gin-gonic/gin"
)

func Apply(c *gin.Context) {
	var data request.Apply
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	uid := c.GetUint("userId")
	register, err := client.Apply(c, &user_enter.UserEnterApplyRequest{
		UeId:         int64(uid),
		Province:     data.Province,
		City:         data.City,
		District:     data.District,
		Address:      data.Address,
		MerchantName: data.MerchantName,
		Charter:      data.Charter,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "申请成功，等待平台审核", register)
}
func Register(c *gin.Context) {
	var data request.Register
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	register, err := client.Register(c, &user_enter.UserEnterRegisterRequest{
		Account:  data.Account,
		Password: data.Password,
		Phone:    data.Phone,
		Email:    data.Email,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "注册成功", register)
}
func Login(c *gin.Context) {
	var data request.Login
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	login, err := client.Login(c, &user_enter.UserEnterLoginRequest{
		Account:  data.Account,
		Password: data.Password,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	claims := pkg.CustomClaims{
		ID: uint(login.UserEnterId),
	}
	token, err := pkg.NewJWT("merchant").CreateToken(claims)
	if err != nil {
		return
	}

	response.RespSuccess(c, 200, "登录成功", token)

}
func AddProduct(c *gin.Context) {
	var data request.AddProduct
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
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
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "发布商品成功", register)
}
func AddCombinationProduct(c *gin.Context) {
	var data request.AddCombinationProduct
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
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
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "发布拼团商品成功", product)
}

// ProcessInvoice TODO:审核发票申请
func ProcessInvoice(c *gin.Context) {
	var data request.ProcessInvoice
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	uid := c.GetUint("userId")
	invoice, err := client.ProcessInvoice(c, &user_enter.ProcessInvoiceRequest{
		UeId:    int64(uid),
		Uid:     data.Uid,
		Status:  data.Status,
		Dis:     data.Dis,
		OrderId: data.OrderId,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}

	if invoice.Greet == false {
		response.RespError(c, 201, "审核失败")
	}

	response.RespSuccess(c, 200, "审核完成", invoice)

}

// UpdateStatus TODO: 下架商品
func UpdateStatus(c *gin.Context) {
	var data request.DelProduct
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	uid := c.GetUint("userId")
	product, err := client.DelProduct(c, &user_enter.DelProductRequest{
		MerId:  int64(uid),
		Pid:    data.Pid,
		Status: data.Status,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	if product.Greet == false {
		response.RespError(c, 201, "下架商品失败")
	}
	response.RespSuccess(c, 200, "下架商品成功", product)

}

// InvoiceList TODO:发票列表展示
func InvoiceList(c *gin.Context) {
	var data request.InvoiceList
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	uid := c.GetUint("userId")
	list, err := client.InvoiceList(c, &user_enter.InvoiceListRequest{
		UeId:   int64(uid),
		Status: data.Status,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "展示成功", list)
}

// BatchPublishProducts TODO: 商品批量发布
func BatchPublishProducts(c *gin.Context) {
	var data request.BatchPublishProducts
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	uid := c.GetUint("userId")
	var productRequests []*user_enter.ProductInfo
	for _, product := range data.Products {
		productRequests = append(productRequests, &user_enter.ProductInfo{
			MerId:     int64(uid),
			Image:     product.Image,
			StoreName: product.StoreName,
			StoreInfo: product.StoreInfo,
			BarCode:   product.BarCode,
			CateId:    product.CateId,
			Price:     float32(product.Price),
			Postage:   float32(product.Postage),
			UnitName:  product.UnitName,
			Activity:  product.Activity,
		})
	}
	list, err := client.BatchPublishProducts(c, &user_enter.BatchPublishProductsRequest{
		Products: productRequests,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "批量发布成功", list)
}

func MerchantVerification(c *gin.Context) {
	var data request.MerchantVerification
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	ar, err := client.MerchantVerification(c, &user_enter.MerchantVerificationRequest{
		UserId:  data.UserId,
		OrderId: data.OrderId,
	})

	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "商家核销成功", ar)
}

func CalculateOrderSummary(c *gin.Context) {
	ar, err := client.CalculateOrderSummary(c, &user_enter.CalculateOrderSummaryRequest{})

	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	response.RespSuccess(c, 200, "商家统计成功", ar)
}

// AddSeckillProduct TODO: 添加秒杀商品
func AddSeckillProduct(c *gin.Context) {
	userEnterId := c.GetUint("userId")
	var data request.AddSeckillProduct
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	seckill, err := client.AddSeckillProduct(c, &user_enter.AddSeckillProductRequest{
		UserEnterId: int64(userEnterId),
		ProductId:   data.ProductId,
		Num:         data.Num,
		Price:       float32(data.Price),
		Description: data.Description,
		StartTime:   data.StartTime,
		StopTime:    data.StopTime,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	if seckill.SeckillId == 0 {
		response.RespError(c, 500, "添加秒杀商品成功")
		return
	}
	response.RespSuccess(c, 200, "添加秒杀商品成功", seckill)
}

// ReverseStock TODO: 秒杀后反还剩余的商品
func ReverseStock(c *gin.Context) {
	userEnterId := c.GetUint("userId")
	var data request.ReverseStock
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	reverse, err := client.ReverseStock(c, &user_enter.ReverseStockRequest{
		UserEnterId: int64(userEnterId),
		SeckillId:   data.ProductId,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	if reverse.Success == false {
		response.RespError(c, 500, "秒杀后反还剩余的商品失败")
		return
	}
	response.RespSuccess(c, 200, "秒杀后反还剩余的商品成功", reverse)
}
