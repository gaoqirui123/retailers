package model

import "common/global"

type OrderCartInfo struct {
	Oid       int64  `gorm:"column:oid;type:int UNSIGNED;comment:订单id;not null;" json:"oid"`                         // 订单id
	CartId    int64  `gorm:"column:cart_id;type:int UNSIGNED;comment:购物车id;not null;default:0;" json:"cart_id"`      // 购物车id
	ProductId int64  `gorm:"column:product_id;type:int UNSIGNED;comment:商品ID;not null;default:0;" json:"product_id"` // 商品ID
	CartInfo  string `gorm:"column:cart_info;type:text;comment:购买东西的详细信息;not null;" json:"cart_info"`                // 购买东西的详细信息
	Unique    string `gorm:"column:unique;type:char(32);comment:唯一id;not null;" json:"unique"`                       // 唯一id
}

func (oc *OrderCartInfo) AddOrderCartInfo() error {
	return global.DB.Debug().Table("order_cart_info").Create(&oc).Error
}
