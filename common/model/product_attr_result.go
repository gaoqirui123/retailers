package model

type ProductAttrResult struct {
	ProductId  uint32 `gorm:"column:product_id;type:int UNSIGNED;comment:商品ID;not null;" json:"product_id"`        // 商品ID
	Result     string `gorm:"column:result;type:text;comment:商品属性参数;not null;" json:"result"`                      // 商品属性参数
	ChangeTime uint32 `gorm:"column:change_time;type:int UNSIGNED;comment:上次修改时间;not null;" json:"change_time"`    // 上次修改时间
	Type       int8   `gorm:"column:type;type:tinyint(1);comment:活动类型 0=商品，1=秒杀，2=砍价，3=拼团;default:0;" json:"type"` // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
}
