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
