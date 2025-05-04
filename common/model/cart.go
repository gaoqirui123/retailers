package model

import (
	"common/global"
	"gorm.io/gorm"
)

type Cart struct {
	Id                uint64 `gorm:"column:id;type:bigint UNSIGNED;comment:购物车表ID;primaryKey;not null;" json:"id"`                  // 购物车表ID
	Uid               uint32 `gorm:"column:uid;type:int UNSIGNED;comment:用户ID;not null;" json:"uid"`                                // 用户ID
	Type              string `gorm:"column:type;type:varchar(32);comment:类型;not null;" json:"type"`                                 // 类型
	ProductId         uint32 `gorm:"column:product_id;type:int UNSIGNED;comment:商品ID;not null;" json:"product_id"`                  // 商品ID
	ProductAttrUnique string `gorm:"column:product_attr_unique;type:varchar(16);comment:商品属性;not null;" json:"product_attr_unique"` // 商品属性
	CartNum           uint16 `gorm:"column:cart_num;type:smallint UNSIGNED;comment:商品数量;not null;default:0;" json:"cart_num"`       // 商品数量
	AddTime           uint32 `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;not null;" json:"add_time"`                      // 添加时间
	IsPay             int8   `gorm:"column:is_pay;type:tinyint(1);comment:0 = 未购买 1 = 已购买;not null;default:0;" json:"is_pay"`       // 0 = 未购买 1 = 已购买
	IsDel             int8   `gorm:"column:is_del;type:tinyint(1);comment:是否删除;not null;default:0;" json:"is_del"`                  // 是否删除
	IsNew             int8   `gorm:"column:is_new;type:tinyint(1);comment:是否为立即购买;not null;default:0;" json:"is_new"`               // 是否为立即购买
	CombinationId     uint32 `gorm:"column:combination_id;type:int UNSIGNED;comment:拼团id;default:0;" json:"combination_id"`         // 拼团id
	SeckillId         uint32 `gorm:"column:seckill_id;type:int UNSIGNED;comment:秒杀商品ID;not null;default:0;" json:"seckill_id"`      // 秒杀商品ID
	BargainId         uint32 `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价id;not null;default:0;" json:"bargain_id"`        // 砍价id
}

func (c *Cart) AddCart() error {
	return global.DB.Debug().Table("cart").Create(&c).Error
}

func (c *Cart) GetStoreCart(uid, productId int64) error {
	return global.DB.Debug().Table("cart").Where("uid = ? and product_id = ?", uid, productId).Limit(1).Find(&c).Error
}

func (c *Cart) UpdateCartNum(uid, productId, num int64) error {
	return global.DB.Debug().Table("cart").Where("uid = ? and product_id = ?", uid, productId).Limit(1).Update("cart_num", gorm.Expr("cart_num + ?", num)).Error
}

func (c *Cart) ClearCart(uid uint64) error {
	return global.DB.Debug().Table("cart").Where("uid = ?", uid).Delete(&c).Error
}

func (c *Cart) DeleteCart(uid, productId int64) error {
	return global.DB.Debug().Table("cart").Where("uid = ? and product_id = ?", uid, productId).Limit(1).Delete(&c).Error
}

func (c *Cart) GetCartList(uid int64) (list []*Cart, err error) {
	err = global.DB.Debug().Table("cart").Where("uid = ?", uid).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Cart) UpdateCart() error {
	return global.DB.Debug().Table("cart").Updates(&c).Error
}
