package model

import (
	"common/global"
	"time"
)

type VisitorCount struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;comment:'自增主键'"`                // ID 自增主键，唯一标识每条访问记录
	ProductID  uint      `gorm:"not null;comment:'商品 ID'"`                               // ProductID 关联的商品 ID，标识被访问的商品
	UserID     uint      `gorm:"not null;comment:'用户 ID'"`                               // UserID 关联的用户 ID，标识访问的用户
	MerchantID uint      `gorm:"not null;comment:'商户 ID'"`                               // MerchantID 关联的商户 ID，标识商品所属的商户
	VisitCount int       `gorm:"default:1;comment:'访问次数'"`                               // VisitCount 记录用户对商品的访问次数，默认值为 1
	CreateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'访问时间'"` // CreateTime 记录访问发生的时间
}

// GetUniqueVisitorCount 获取商品访问客数，假设访客表中有一个 user_id 字段来区分不同访客
func (v *VisitorCount) GetUniqueVisitorCount() (int64, error) {
	var count int64
	err := global.DB.Model(v).Distinct("user_id").Count(&count).Error
	return count, err
}
