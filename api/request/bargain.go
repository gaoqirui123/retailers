package request

// 创建砍价商品信息
type BargainCreate struct {
	UserID          uint32  `json:"user_id" form:"user_id" binding:"required"`                     // 必填，标识发起砍价的用户 ID
	ProductId       uint32  `json:"product_id" form:"product_id" binding:"required"`               // 与砍价活动关联的商品 ID
	Title           string  `json:"title" form:"title" binding:"required"`                         // 砍价活动的名称
	Image           string  `json:"image" form:"image" binding:"required"`                         // 砍价活动对应的图片地址
	UnitName        string  `json:"unit_name" form:"unit_name" binding:"required"`                 // 商品的单位名称
	Stock           uint32  `json:"stock" form:"stock" binding:"required"`                         // 参与砍价商品的库存数量
	Images          string  `json:"images" form:"images" binding:"required"`                       // 砍价商品的轮播图信息（如图片地址列表等）
	Price           float64 `json:"price" form:"price" binding:"required"`                         // 砍价金额，对应数据库 decimal 类型，使用 double 保证精度
	MinPrice        float64 `json:"min_price" form:"min_price" binding:"required"`                 // 砍价商品的最低价格，对应数据库 decimal 类型，使用 double 保证精度
	Num             uint32  `json:"num" form:"num" binding:"required"`                             // 每次购买砍价商品的数量
	BargainMaxPrice float64 `json:"bargain_max_price" form:"bargain_max_price" binding:"required"` // 用户每次砍价可达到的最大金额，对应数据库 decimal 类型，使用 double 保证精度
	BargainMinPrice float64 `json:"bargain_min_price" form:"bargain_min_price" binding:"required"` // 用户每次砍价可达到的最小金额，对应数据库 decimal 类型，使用 double 保证精度
	BargainNum      uint32  `json:"bargain_num" form:"bargain_num" binding:"required"`             // 用户每次砍价的次数
	Status          uint32  `json:"status" form:"status" binding:"required"`                       // 砍价状态，0 表示到砍价时间不自动开启，1 表示到砍价时间自动开启
	GiveIntegral    float64 `json:"give_integral" form:"give_integral" binding:"required"`         // 参与砍价成功后返还的积分数量，对应数据库 decimal 类型，使用 double 保证精度
	Info            string  `json:"info" form:"info" binding:"required"`                           // 砍价活动的详细介绍信息
	IsPostage       uint32  `json:"is_postage" form:"is_postage" binding:"required"`               // 是否包邮标识，0 为不包邮，1 为包邮
	Postage         float64 `json:"postage" form:"postage" binding:"required"`                     // 商品的邮费金额，对应数据库 decimal 类型，使用 double 保证精度
	Rule            string  `json:"rule" form:"rule" binding:"required"`                           // 砍价活动的具体规则说明
	StoreName       string  `json:"store_name" form:"store_name" binding:"required"`               // 砍价商品所属店铺或商家名称
	TempId          int32   `json:"temp_id" form:"temp_id" binding:"required"`                     // 运费模板 ID
	Cost            float64 `json:"cost" form:"cost" binding:"required"`                           //成本价
}

type ProductUpdate struct {
	Id        uint32 `json:"id" form:"id" binding:"required"` // 砍价表的唯一标识 ID
	IsBargain int32  `json:"is_bargain" form:"is_bargain"`    //是否砍价
}
