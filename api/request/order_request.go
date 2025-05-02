package request

type AddOrder struct {
	ProductId             int64  `json:"productId" form:"productId" binding:"required"`
	Num                   int64  `json:"num" form:"num" binding:"required"`
	PayType               int64  `json:"payType" form:"payType"`
	CouponId              int64  `json:"couponId" form:"couponId"`
	Mark                  string `json:"mark" form:"mark"`
	StoreId               int64  `json:"storeId" form:"storeId"`
	MerId                 int64  `json:"merId" form:"merId"`
	BargainId             int64  `json:"bargainId" form:"bargainId"`
	ShippingType          int64  `json:"shippingType" form:"shippingType"`
	IsChannel             int64  `json:"isChannel" form:"isChannel"`
	PinkId                int64  `json:"pinkId" form:"pinkId"`
	ProductSpecifications string `json:"productSpecifications" form:"productSpecifications" binding:"required"`
}
type OrderList struct {
	OrderStatus int64 `json:"orderStatus" form:"orderStatus"`
}
