package model

type ProductCate struct {
	Id        int32 `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	ProductId int32 `gorm:"column:product_id;type:int;comment:商品id;not null;default:0;" json:"product_id"` // 商品id
	CateId    int32 `gorm:"column:cate_id;type:int;comment:分类id;not null;default:0;" json:"cate_id"`       // 分类id
	AddTime   int32 `gorm:"column:add_time;type:int;comment:添加时间;not null;default:0;" json:"add_time"`     // 添加时间
}
