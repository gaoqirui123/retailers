package request

type Apply struct {
	Province     string `json:"province" xml:"province" form:"province" binding:"required"`
	City         string `json:"city" xml:"city" form:"city" binding:"required"`
	District     string `json:"district" xml:"district" form:"district" binding:"required"`
	Address      string `json:"address" xml:"address" form:"address" binding:"required"`
	MerchantName string `json:"merchantName" xml:"merchantName" form:"merchantName" binding:"required"`
	LinkTel      string `json:"linkTel" xml:"linkTel" form:"linkTel" binding:"required"`
	Charter      string `json:"charter" xml:"charter" form:"charter" binding:"required"`
}
type Register struct {
	Account  string `json:"account" xml:"account" form:"account" binding:"required"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
	Phone    string `json:"phone" xml:"phone" form:"phone" binding:"required"`
	Email    string `json:"email" xml:"email" form:"email" binding:"required"`
}
type Login struct {
	Account  string `json:"account" xml:"account" form:"account" binding:"required"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}
type AddProduct struct {
	Image     string  `json:"image" xml:"image" form:"image" binding:"required"`
	StoreName string  `json:"storeName" xml:"storeName" form:"storeName" binding:"required"`
	StoreInfo string  `json:"storeInfo" xml:"storeInfo" form:"storeInfo" binding:"required"`
	BarCode   string  `json:"barCode" xml:"barCode" form:"barCode" binding:"required"`
	CateId    string  `json:"cateId" xml:"cateId" form:"cateId" binding:"required"`
	Price     float64 `json:"price" xml:"price" form:"price" binding:"required"`
	Postage   float64 `json:"postage" xml:"postage" form:"postage" binding:"required"`
	UnitName  string  `json:"unitName" xml:"unitName" form:"unitName" binding:"required"`
	Activity  string  `json:"activity" xml:"activity" form:"activity" binding:"required"`
}
type AddCombinationProduct struct {
	ProductId     int64   `json:"productId" xml:"productId" form:"productId"`
	Title         string  `json:"title" xml:"title" form:"title"`
	Attr          string  `json:"attr" xml:"attr" form:"attr"`
	People        int64   `json:"people" xml:"people" form:"people"`
	Price         float64 `json:"price" xml:"price" form:"price"`
	Sort          int32   `json:"sort" xml:"sort" form:"sort"`
	Stock         int32   `json:"stock" xml:"stock" form:"stock"`
	StartTime     int32   `json:"startTime" xml:"startTime" form:"startTime"`
	StopTime      int32   `json:"stopTime" xml:"stopTime" form:"stopTime"`
	EffectiveTime int32   `json:"effectiveTime" xml:"effectiveTime" form:"effectiveTime"`
	TempId        int32   `json:"tempId" xml:"tempId" form:"tempId"`
	Num           int32   `json:"num" xml:"num" form:"num"`
	Quota         int32   `json:"quota" xml:"quota" form:"quota"`
	QuotaShow     int32   `json:"quotaShow" xml:"quotaShow" form:"quotaShow"`
}
type ProcessInvoice struct {
	Status  int64  `json:"status" xml:"status" form:"status"`
	Uid     int64  `json:"uid" xml:"uid" form:"uid"`
	Dis     string `json:"dis" xml:"dis" form:"dis"`
	OrderId int64  `json:"orderId" xml:"orderId" form:"orderId"`
}
type DelProduct struct {
	Pid    int64 `json:"pid" xml:"pid" form:"pid"`
	Status int64 `json:"status" xml:"status" form:"status"`
}
type InvoiceList struct {
	Status int64 `json:"status" xml:"status" form:"status"`
}

type AddSeckillProduct struct {
	ProductId   int64   `json:"productId" xml:"productId" form:"productId"`
	Num         int64   `json:"num" xml:"num" form:"num"`
	Price       float64 `json:"price" xml:"price" form:"price"`
	Description string  `json:"description" xml:"description" form:"description"`
	StartTime   string  `json:"startTime" xml:"startTime" form:"startTime"`
	StopTime    string  `json:"stopTime" xml:"stopTime" form:"stopTime"`
}

type ReverseStock struct {
	ProductId int64 `json:"productId" xml:"productId" form:"productId"`
}
