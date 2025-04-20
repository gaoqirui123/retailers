package model

import "common/global"

type OrderProduct struct {
	Id                    uint32  `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	OrderId               uint32  `gorm:"column:order_id;type:int UNSIGNED;comment:订单ID;not null;" json:"order_id"`                            // 订单ID
	ProductId             uint32  `gorm:"column:product_id;type:int UNSIGNED;comment:商品Id;not null;" json:"product_id"`                        // 商品Id
	ProductName           string  `gorm:"column:product_name;type:varchar(50);comment:商品名称;not null;" json:"product_name"`                     // 商品名称
	ProductImage          string  `gorm:"column:product_image;type:varchar(255);comment:商品图频;not null;" json:"product_image"`                  // 商品图频
	ProductSpecifications string  `gorm:"column:product_specifications;type:varchar(50);comment:商品规格;not null;" json:"product_specifications"` // 商品规格
	ProductPrice          float64 `gorm:"column:product_price;type:decimal(10, 2);comment:商品价格;not null;" json:"product_price"`                // 商品价格
	ProductNum            int32   `gorm:"column:product_num;type:int;comment:商品数量;not null;" json:"product_num"`                               // 商品数量
}

func (p *OrderProduct) GetOrderProductIdBy(orderId int64) error {
	return global.DB.Debug().Table("order_product").Where("order_id = ?", orderId).Limit(1).Find(&p).Error
}
