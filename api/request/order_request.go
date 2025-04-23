package request

type AddOrder struct {
	ProductId     int64  `json:"productId" form:"productId" binding:"required"`
	Num           int64  `json:"num" form:"num" binding:"required"`
	PayType       int64  `json:"payType" form:"payType" binding:"required"`
	CartId        int64  `json:"cartId" form:"cartId"`
	FreightPrice  int64  `json:"freightPrice" form:"freightPrice" binding:"required"`
	CouponId      int64  `json:"couponId" form:"couponId" binding:"required"`
	UseIntegral   int64  `json:"useIntegral" form:"useIntegral" binding:"required"`
	Mark          string `json:"mark" form:"mark"`
	StoreId       int64  `json:"storeId" form:"storeId" binding:"required"`
	MerId         int64  `json:"merId" form:"merId" binding:"required"`
	BargainId     int64  `json:"bargainId" form:"bargainId" binding:"required"`
	ShippingType  int64  `json:"shippingType" form:"shippingType"`
	IsChannel     int64  `json:"isChannel" form:"isChannel"`
	CombinationId int64  `json:"combinationId" form:"combinationId"`
	PinkId        int64  `json:"pinkId" form:"pinkId"`
}
type OrderList struct {
	OrderStatus int64 `json:"orderStatus" form:"orderStatus"`
}
