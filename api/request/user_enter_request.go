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
}
