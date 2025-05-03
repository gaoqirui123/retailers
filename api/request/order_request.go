package request

type AddOrder struct {
	ProductId             int64  `json:"productId" form:"productId" binding:"required"`
	Num                   int64  `json:"num" form:"num" binding:"required"`
	PayType               int64  `json:"payType" form:"payType" binding:"required"`
	CouponId              int64  `json:"couponId" form:"couponId"`
	Mark                  string `json:"mark" form:"mark"`
	StoreId               int64  `json:"storeId" form:"storeId" binding:"required"`
	MerId                 int64  `json:"merId" form:"merId" binding:"required"`
	BargainId             int64  `json:"bargainId" form:"bargainId"`
	ShippingType          int64  `json:"shippingType" form:"shippingType" binding:"required"`
	IsChannel             int64  `json:"isChannel" form:"isChannel" binding:"required"`
	PinkId                int64  `json:"pinkId" form:"pinkId"`
	ProductSpecifications string `json:"productSpecifications" form:"productSpecifications" binding:"required"`
	Source                int64  `json:"source" form:"source"`
}
type OrderList struct {
	OrderStatus int64 `json:"orderStatus" form:"orderStatus"`
}
