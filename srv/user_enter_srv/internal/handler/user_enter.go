package handler

import (
	"common/global"
	"common/model"
	"common/pkg"
	"common/proto/user_enter"
	"common/utlis"
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"
)

type OrderSyn struct {
	ID      int64  `json:"id"`
	OrderSn string `json:"orderSn"`
	Status  int64  `json:"status"`
	Paid    int64  `json:"paid"`
}

// Apply TODO:商家申请店铺
func Apply(in *user_enter.UserEnterApplyRequest) (*user_enter.UserEnterApplyResponse, error) {
	ue := model.UserEnter{
		Uid:          int(in.UeId),
		Province:     in.Province,
		City:         in.City,
		District:     in.District,
		Address:      in.Address,
		MerchantName: in.MerchantName,
		LinkTel:      in.LinkTel,
		Charter:      in.Charter,
	}
	if in.MerchantName == "" {
		return nil, errors.New("商户名称不能为空")
	}
	if in.LinkTel == "" {
		return nil, errors.New("商户电话不能为空")
	}
	err := ue.Add()
	if err != nil {
		return nil, err
	}
	return &user_enter.UserEnterApplyResponse{Greet: true}, nil
}

// AddProduct TODO:商家发布商品
func AddProduct(in *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
	p := model.Product{
		MerId:     in.MerId,
		Image:     in.Image,
		StoreName: in.StoreName,
		StoreInfo: in.StoreInfo,
		BarCode:   in.BarCode,
		CateId:    in.CateId,
		Price:     float64(in.Price),
		Postage:   float64(in.Postage),
		UnitName:  in.UnitName,
		Activity:  in.Activity,
	}
	err := p.Add()
	if err != nil {
		return nil, err
	}
	return &user_enter.AddProductResponse{Greet: true}, nil
}

// AddCombinationProduct TODO:商家发布拼团商品
func AddCombinationProduct(in *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
	p := model.Product{}
	product, err := p.GetProductById(in.ProductId, 0)
	if err != nil {
		return nil, err
	}
	c := &model.Combination{
		MerId:         int(in.MerId),
		ProductId:     int(in.ProductId),
		Image:         product.Image,
		Images:        product.Image,
		Title:         in.Title,
		Attr:          in.Attr,
		People:        int(in.People),
		Info:          product.StoreInfo,
		Price:         float64(in.Price),
		Sort:          int(in.Sort),
		Stock:         int(in.Stock),
		StartTime:     int(in.StartTime),
		StopTime:      int(in.StopTime),
		EffectiveTime: in.EffectiveTime,
		Cost:          int(product.Cost),
		UnitName:      product.UnitName,
		TempId:        in.TempId,
		Num:           in.Num,
		Quota:         in.Quota,
		QuotaShow:     in.QuotaShow,
	}
	err = c.Add()
	if err != nil {
		return nil, err
	}
	return &user_enter.AddCombinationProductResponse{Greet: true}, nil
}

// ProcessInvoice TODO:商家审核用户的发票invoice申请
func ProcessInvoice(in *user_enter.ProcessInvoiceRequest) (*user_enter.ProcessInvoiceResponse, error) {
	// 检查输入参数是否为空
	if in == nil {
		return nil, errors.New("请求参数有误")
	}
	//查找用户是否有申请
	i := model.InvoiceApplication{}
	invoiceByUeId, err := i.GetInvoiceByUeId(in.Uid, in.OrderId)
	if err != nil {
		return nil, err
	}
	// 查找用户订单
	od := model.Order{}
	_, err = od.FindUserOrder(in.Uid, invoiceByUeId.OrderId)
	if err != nil {
		return nil, err
	}

	// 更新发票申请状态
	i = model.InvoiceApplication{}
	if in.Dis == "" {
		err = i.UpdateStatus(in.UeId, in.Uid, in.Status, in.OrderId, time.Now())
		if err != nil {
			return nil, errors.New("审核失败")
		}
	} else {
		err = i.UpdateStatusDis(in.UeId, in.Uid, in.Status, in.Dis, in.OrderId, time.Now())
		if err != nil {
			return nil, errors.New("审核失败")
		}
	}
	id, _ := i.GetInvoiceByUeId(in.Uid, in.OrderId)
	//发票审核成功后生成发票图片
	err = pkg.GenerateInvoiceImage(id)
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Fatalf("生成图片时出错: %v", err)
	}
	return &user_enter.ProcessInvoiceResponse{Greet: true}, nil
}

// InvoiceList TODO:发票列表展示
func InvoiceList(in *user_enter.InvoiceListRequest) (*user_enter.InvoiceListResponse, error) {
	ue := model.InvoiceApplication{}
	var lists []*user_enter.InvoiceList
	var applications []*model.InvoiceApplication
	var err error
	if in.Status == 0 {
		applications, err = ue.GetInvoiceByUeIds(in.UeId)
		if err != nil {
			return nil, errors.New("没有该商户的发票资料")
		}
	} else {
		applications, err = ue.GetInvoiceByUeIdAndStatus(in.UeId, in.Status)
		if err != nil {
			return nil, err
		}
	}

	for _, application := range applications {
		list := ue.ConvertToInvoiceList(application)
		lists = append(lists, &list)
	}

	return &user_enter.InvoiceListResponse{
		List: lists,
	}, nil
}

// DelProduct TODO:下架商品
func DelProduct(in *user_enter.DelProductRequest) (*user_enter.DelProductResponse, error) {
	p := model.Product{}
	product, err := p.GetProductById(in.MerId, in.Pid)
	if err != nil {
		return nil, err
	}
	if product.Id == 0 {
		return nil, errors.New("没有该商品")
	}
	//判断商品是否上架
	if product.IsShow == 0 {
		return nil, errors.New("该商品未上架")
	}
	err = product.UpdateStatus(in.Status, in.Pid)
	if err != nil {
		return nil, err
	}
	return &user_enter.DelProductResponse{Greet: true}, nil
}

// Register TODO:商家注册账号
func Register(in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	m := model.Merchant{
		MerchantAccount:  in.Account,
		MerchantPassword: in.Password,
		ContactPhone:     in.Phone,
		Email:            in.Email,
	}
	merchant, err := m.GetMerchantByAccount(in.Account)
	if merchant.MerchantId != 0 {
		return nil, errors.New("该账号已注册")
	}
	regex := `^1[3-9]\d{9}$`
	matchString, err := regexp.MatchString(regex, in.Phone)
	if !matchString {
		return nil, errors.New("手机号码不规范")
	}
	merchant, err = m.GetMerchantByPhone(in.Phone)
	if merchant.MerchantId != 0 {
		return nil, errors.New("该手机号已注册")
	}
	merchant, err = m.GetMerchantByEmail(in.Email)
	if merchant.MerchantId != 0 {
		return nil, errors.New("该邮箱已注册")
	}
	err = m.Create()
	if err != nil {
		return nil, errors.New("注册失败")
	}
	return &user_enter.UserEnterRegisterResponse{Greet: true}, nil
}

// Login TODO:商家账号登录
func Login(in *user_enter.UserEnterLoginRequest) (*user_enter.UserEnterLoginResponse, error) {
	m := model.Merchant{
		MerchantAccount:  in.Account,
		MerchantPassword: in.Password,
	}
	merchant, _ := m.GetMerchantByAccount(in.Account)
	if merchant.MerchantId == 0 {
		return nil, errors.New("该账号未注册")
	}
	if in.Password != merchant.MerchantPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	claims := pkg.CustomClaims{
		ID: uint(merchant.MerchantId),
	}
	token, err := pkg.NewJWT("merchant").CreateToken(claims)
	if err != nil {
		return nil, err
	}
	return &user_enter.UserEnterLoginResponse{Greet: token}, nil
}

// AddSeckillProduct  TODO: 添加秒杀商品
func AddSeckillProduct(in *user_enter.AddSeckillProductRequest) (*user_enter.AddSeckillProductResponse, error) {
	//查询商品是否存在
	p := &model.Product{}
	err := p.GetProductIdBy(in.ProductId)
	if err != nil {
		return nil, err
	}
	if p.Id == 0 {
		return nil, errors.New("商品不存在")
	}
	// 判断该商品是否是该商户的
	if p.MerId != in.UserEnterId {
		return nil, errors.New("该商品不是你的")
	}
	//判断商品库存不能小于秒杀库存
	if p.Stock < in.Num {
		return nil, errors.New("判断商品库存小于秒杀库存")
	}
	//开启事务
	tx := global.DB.Begin()
	//扣mysql商品表总库存
	err = p.UpdateProductStock(in.ProductId, in.Num)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("扣mysql商品表总库存失败")
	}
	seckill := &model.Seckill{
		MerId:       p.MerId,
		ProductId:   in.ProductId,
		Image:       p.Image,
		Images:      p.SliderImage,
		Name:        p.StoreName,
		Info:        p.StoreInfo,
		Price:       float64(in.Price),
		Cost:        p.Cost,
		OtPrice:     p.Price,
		Stock:       in.Num,
		Postage:     p.Postage,
		Description: in.Description,
		StartTime:   in.StartTime,
		StopTime:    in.StopTime,
		AddTime:     time.Now().Format(time.DateTime),
		Status:      p.IsShow,
		IsPostage:   p.IsPostage,
		Num:         in.Num,
		Quota:       in.Num,
		QuotaShow:   in.Num,
	}
	err = seckill.AddSeckillProduct()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if seckill.Id == 0 {
		tx.Rollback()
		return nil, errors.New("添加秒杀商品失败")
	}
	//将秒杀商品添加redis的list中
	utlis.SeckillCreateRedis(int(seckill.Stock), int(seckill.Id))
	//判断redis库存是否添加成功
	val := utlis.GetSeckillRedis(int(seckill.Id))
	if val != seckill.Stock {
		tx.Rollback()
		return nil, errors.New("redis库存是否添加失败")
	}
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	return &user_enter.AddSeckillProductResponse{SeckillId: seckill.Id}, nil
}

// ReverseStock  TODO: 秒杀后反还剩余的商品
func ReverseStock(in *user_enter.ReverseStockRequest) (*user_enter.ReverseStockResponse, error) {
	// 查询秒杀商品是否存在
	s := &model.Seckill{}
	err := s.GetSeckillIdBY(in.SeckillId)
	if err != nil {
		return nil, err
	}
	if s.Id == 0 {
		return nil, errors.New("秒杀商品不存在")
	}
	// 判断该商品是否是该商户的
	if s.MerId != in.UserEnterId {
		return nil, errors.New("该商品不是你的")
	}
	//开启事务
	tx := global.DB.Begin()
	//反还剩余的商品
	g := &model.Product{}
	err = g.ReverseProductStock(s.ProductId, s.Stock)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("秒杀商品不存在")
	}
	//清除秒杀表里的数据
	err = s.DelSeckill(in.SeckillId)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("清除秒杀表里的数据失败")
	}
	//清除redis列表库存
	utlis.DelSeckillRedis(int(s.Id))
	//判断redis列表库存是否被清除
	val := utlis.GetSeckillRedis(int(s.Id))
	if val > 0 {
		tx.Rollback()
		return nil, errors.New("清除redis列表库存失败")
	}
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	return &user_enter.ReverseStockResponse{Success: true}, nil
}

// BatchReleaseOfProducts TODO:商品批量发布
func BatchReleaseOfProducts(in *user_enter.BatchReleaseOfProductsRequest) (*user_enter.BatchReleaseOfProductsResponse, error) {
	// 开启事务
	transaction := global.DB.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	for _, productData := range in.List {
		p := model.Product{
			MerId:       productData.MerId,
			Image:       productData.Image,
			SliderImage: productData.SliderImage,
			StoreName:   productData.StoreName,
			CateId:      strconv.FormatInt(productData.CateId, 10),
			IsShow:      int64(int(productData.IsShow)),
			Price:       float64(productData.Price),
			Postage:     float64(productData.Postage),
			UnitName:    productData.UnitName,
		}

		// 在事务中添加商品
		if err := p.Add(); err != nil {
			// 回滚事务
			transaction.Rollback()
			return nil, err
		}
	}

	// 提交事务
	if err := transaction.Commit().Error; err != nil {
		return nil, err
	}

	return &user_enter.BatchReleaseOfProductsResponse{Success: "批量添加商品成功"}, nil
}

// MerchantVerification TODO:商家核销
func MerchantVerification(in *user_enter.MerchantVerificationRequest) (*user_enter.MerchantVerificationResponse, error) {
	o := &model.Order{}
	id, err := o.FindId(in.OrderId)
	if err != nil {
		return &user_enter.MerchantVerificationResponse{Greet: false}, nil
	}
	err = o.UpdateOrderStatus(id.OrderSn, 7)
	if err != nil {
		return &user_enter.MerchantVerificationResponse{Greet: false}, nil
	}
	return &user_enter.MerchantVerificationResponse{Greet: true}, nil
}

// CalculateOrderSummary TODO:商家统计
func CalculateOrderSummary(in *user_enter.CalculateOrderSummaryRequest) (*user_enter.CalculateOrderSummaryResponse, error) {
	orders := &model.Order{}
	products := &model.Product{}
	// 获取订单总数
	orderCount, err := orders.GetTotalOrderCount()
	if err != nil {
		return nil, err
	}

	// 获取订单总金额
	totalAmount, err := orders.GetTotalOrderAmount()
	if err != nil {
		return nil, err
	}

	// 获取总退款数
	totalRefund, err := orders.GetTotalRefundAmount()
	if err != nil {
		return nil, err
	}

	// 获取商品总浏览量
	productViewCount, err := products.GetTotalViewCount()
	if err != nil {
		return nil, err
	}

	// 获取商品访问客数
	//uniqueVisitors, err := visitors.GetUniqueVisitorCount()
	//if err != nil {
	//	return nil, err
	//}

	return &user_enter.CalculateOrderSummaryResponse{
		OrderCount:       int32(orderCount),
		TotalAmount:      float32(totalAmount),
		TotalRefund:      float32(totalRefund),
		ProductViewCount: int32(productViewCount),
		//	UniqueVisitors:   int32(uniqueVisitors),
	}, nil
}
