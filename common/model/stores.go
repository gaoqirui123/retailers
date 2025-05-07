package model

import (
	"common/global"
	"time"
)

// Stores 商家店铺表
type Stores struct {
	StoreId          int64     `gorm:"column:store_id;type:int;primaryKey;not null;" json:"store_id"`
	MerchantId       int64     `gorm:"column:merchant_id;type:int;comment:商户id;not null;" json:"merchant_id"`                     // 商户id
	StoreName        string    `gorm:"column:store_name;type:varchar(255);comment:店铺名称;not null;" json:"store_name"`              // 店铺名称
	StoreDescription string    `gorm:"column:store_description;type:text;comment:店铺介绍;" json:"store_description"`                 // 店铺介绍
	StoreLogo        string    `gorm:"column:store_logo;type:varchar(255);comment:店铺logo;default:NULL;" json:"store_logo"`        // 店铺logo
	ContactPhone     string    `gorm:"column:contact_phone;type:varchar(20);comment:电话;default:NULL;" json:"contact_phone"`       // 电话
	Address          string    `gorm:"column:address;type:varchar(255);comment:地址;default:NULL;" json:"address"`                  // 地址
	Rating           float64   `gorm:"column:rating;type:decimal(3, 2);comment:店铺评分;default:0.00;" json:"rating"`                 // 店铺评分
	ReviewCount      int64     `gorm:"column:review_count;type:int;comment:评论数量;default:0;" json:"review_count"`                  // 评论数量
	IsActive         int64     `gorm:"column:is_active;type:tinyint(1);comment:是否活跃;default:1;" json:"is_active"`                 // 是否活跃
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 修改时间
}

func (s *Stores) TableName() string {
	return "stores"
}
func (s *Stores) CreateStores() error {
	return global.DB.Create(&s).Error
}
