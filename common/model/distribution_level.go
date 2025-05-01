package model

import (
	"common/global"
	"time"
)

// 分销等级表
type DistributionLevel struct {
	Id          uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Img         string    `gorm:"column:img;type:varchar(255);comment:商品名称;not null;" json:"img"`                  // 商品名称
	LevelName   string    `gorm:"column:level_name;type:varchar(255);comment:等级名称;not null;" json:"level_name"`    // 等级名称
	Level       int64     `gorm:"column:level;type:bigint;comment:1/2/3/;not null;default:1;" json:"level"`        // 1/2/3/
	One         float64   `gorm:"column:one;type:decimal(10, 2);comment:一级返佣比例;not null;default:0.00;" json:"one"` // 一级返佣比例
	Two         float64   `gorm:"column:two;type:decimal(10, 2);comment:二级返佣比例;not null;default:0.00;" json:"two"` // 二级返佣比例
	Start       int64     `gorm:"column:start;type:bigint;comment:1开启/2未开启;not null;default:1;" json:"start"`      // 1开启/2未开启
	CreatedTime time.Time `gorm:"column:created_time;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"created_time"`
	IsDel       int64     `gorm:"column:is_del;type:bigint;comment:1未删除/2已删;not null;default:1;" json:"is_del"` // 1未删除/2已删
}

func (dl *DistributionLevel) Table() string {
	return "distribution_level"
}
func (dl *DistributionLevel) CreateDistributionLevel() bool {
	err := global.DB.Debug().Table("distribution_level").Create(&dl).Error
	if err != nil {
		return false
	}
	return true
}

func (dl *DistributionLevel) FindDistributionLevel(level int) DistributionLevel {
	var distribution DistributionLevel
	err := global.DB.Debug().Table("distribution_level").Where("level = ?", level).Find(&distribution).Error
	if err != nil {
		return distribution
	}
	return distribution
}
