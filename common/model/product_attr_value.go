package model

type ProductAttrValue struct {
	ProductId    uint32  `gorm:"column:product_id;type:int UNSIGNED;comment:商品ID;not null;" json:"product_id"`                     // 商品ID
	Suk          string  `gorm:"column:suk;type:varchar(128);comment:商品属性索引值 (attr_value|attr_value[|....]);not null;" json:"suk"` // 商品属性索引值 (attr_value|attr_value[|....])
	Stock        uint32  `gorm:"column:stock;type:int UNSIGNED;comment:属性对应的库存;not null;" json:"stock"`                            // 属性对应的库存
	Sales        uint32  `gorm:"column:sales;type:int UNSIGNED;comment:销量;not null;default:0;" json:"sales"`                       // 销量
	Price        float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:属性金额;not null;" json:"price"`                     // 属性金额
	Image        string  `gorm:"column:image;type:varchar(128);comment:图片;default:NULL;" json:"image"`                             // 图片
	Unique       string  `gorm:"column:unique;type:char(8);comment:唯一值;not null;" json:"unique"`                                   // 唯一值
	Cost         float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本价;not null;" json:"cost"`                        // 成本价
	BarCode      string  `gorm:"column:bar_code;type:varchar(50);comment:商品条码;not null;" json:"bar_code"`                          // 商品条码
	OtPrice      float64 `gorm:"column:ot_price;type:decimal(8, 2);comment:原价;not null;default:0.00;" json:"ot_price"`             // 原价
	Weight       float64 `gorm:"column:weight;type:decimal(8, 2);comment:重量;not null;default:0.00;" json:"weight"`                 // 重量
	Volume       float64 `gorm:"column:volume;type:decimal(8, 2);comment:体积;not null;default:0.00;" json:"volume"`                 // 体积
	Brokerage    float64 `gorm:"column:brokerage;type:decimal(8, 2);comment:一级返佣;not null;default:0.00;" json:"brokerage"`         // 一级返佣
	BrokerageTwo float64 `gorm:"column:brokerage_two;type:decimal(8, 2);comment:二级返佣;not null;default:0.00;" json:"brokerage_two"` // 二级返佣
	Type         int8    `gorm:"column:type;type:tinyint(1);comment:活动类型 0=商品，1=秒杀，2=砍价，3=拼团;default:0;" json:"type"`              // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
	Quota        int32   `gorm:"column:quota;type:int;comment:活动限购数量;default:NULL;" json:"quota"`                                  // 活动限购数量
	QuotaShow    int32   `gorm:"column:quota_show;type:int;comment:活动限购数量显示;default:NULL;" json:"quota_show"`                      // 活动限购数量显示
}
