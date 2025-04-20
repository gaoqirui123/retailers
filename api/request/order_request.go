package request

type AddOrder struct {
	ProductId     int64  `json:"productId" form:"productId" binding:"required"`
	Num           int64  `json:"num" form:"num" binding:"required"`
	PayType       int64  `json:"payType" form:"passWord" binding:"required"`
	CartId        int64  `json:"cartId" form:"passWord" binding:"required"`
	FreightPrice  int64  `json:"freightPrice" form:"passWord" binding:"required"`
	CouponId      int64  `json:"couponId" form:"passWord" binding:"required"`
	UseIntegral   int64  `json:"useIntegral" form:"passWord" binding:"required"`
	Mark          string `json:"mark" form:"passWord" binding:"required"`
	StoreId       int64  `json:"storeId" form:"passWord" binding:"required"`
	MerId         int64  `json:"merId" form:"passWord" binding:"required"`
	BargainId     int64  `json:"bargainId" form:"passWord" binding:"required"`
	ShippingType  int64  `json:"shippingType" form:"passWord" binding:"required"`
	IsChannel     int64  `json:"isChannel" form:"passWord" binding:"required"`
	CombinationId int64  `json:"combinationId" form:"passWord" binding:"required"`
	PinkId        int64  `json:"pinkId" form:"passWord" binding:"required"`
}
type OrderList struct {
	OrderStatus int64 `json:"orderStatus" form:"productId" binding:"required"`
}
