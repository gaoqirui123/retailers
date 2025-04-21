package request

type Register struct {
	Province     string `json:"province" xml:"province" form:"province"`
	City         string `json:"city" xml:"city" form:"city"`
	District     string `json:"district" xml:"district" form:"district"`
	Address      string `json:"address" xml:"address" form:"address"`
	MerchantName string `json:"merchantName" xml:"merchantName" form:"merchantName"`
	LinkTel      string `json:"linkTel" xml:"linkTel" form:"linkTel"`
	Charter      string `json:"charter" xml:"charter" form:"charter"`
}
type AddProduct struct {
	MerId     int64   `json:"merId" xml:"merId" form:"merId"`
	Image     string  `json:"image" xml:"image" form:"image"`
	StoreName string  `json:"storeName" xml:"storeName" form:"storeName"`
	StoreInfo string  `json:"storeInfo" xml:"storeInfo" form:"storeInfo"`
	BarCode   string  `json:"barCode" xml:"barCode" form:"barCode"`
	CateId    string  `json:"cateId" xml:"cateId" form:"cateId"`
	Price     float64 `json:"price" xml:"price" form:"price"`
	Postage   float64 `json:"postage" xml:"postage" form:"postage"`
	UnitName  string  `json:"unitName" xml:"unitName" form:"unitName"`
	Activity  string  `json:"activity" xml:"activity" form:"activity"`
}
