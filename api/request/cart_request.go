package request

type AddCart struct {
	ProductId         int64  `json:"productId" form:"productId" binding:"required"`
	Type              string `json:"type" form:"type" binding:"required"`
	ProductAttrUnique string `json:"productAttrUnique" form:"productAttrUnique" binding:"required"`
	CartNum           int64  `json:"cartNum" form:"cartNum" binding:"required"`
	IsPay             int64  `json:"isPay" form:"isPay"`
	IsNew             int64  `json:"isNew" form:"isNew"`
	CombinationId     int64  `json:"combinationId" form:"combinationId"`
	SeckillId         int64  `json:"seckillId" form:"seckillId"`
	BargainId         int64  `json:"bargainId" form:"bargainId"`
}

type DeleteCart struct {
	ProductId int64 `json:"productId" form:"productId" binding:"required"`
}
