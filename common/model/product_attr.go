package model

type ProductAttr struct {
	ProductId  uint32 `gorm:"column:product_id;type:int UNSIGNED;comment:商品ID;not null;default:0;" json:"product_id"` // 商品ID
	AttrName   string `gorm:"column:attr_name;type:varchar(32);comment:属性名;not null;" json:"attr_name"`               // 属性名
	AttrValues string `gorm:"column:attr_values;type:varchar(256);comment:属性值;not null;" json:"attr_values"`          // 属性值
	Type       int8   `gorm:"column:type;type:tinyint(1);comment:活动类型 0=商品，1=秒杀，2=砍价，3=拼团;default:0;" json:"type"`    // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
}
