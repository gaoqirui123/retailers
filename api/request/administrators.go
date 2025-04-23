package request

type AdminLogin struct {
	Account string `json:"account" xml:"account" form:"account"`
	Pwd     string `json:"pwd" xml:"pwd" form:"pwd"`
}
type ProcessEnter struct {
	MerchantId int64 `json:"merchantId" xml:"merchantId" form:"merchantId"` //商户id
	Status     int64 `json:"status" xml:"status" form:"status"`             //审核状态
}
