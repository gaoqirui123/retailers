package user_level

import (
	"common/global"
	"time"
)

// 会员分添加记录表
type UserLevelAddText struct {
	Id      uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Uid     uint32    `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                                 // 用户id
	Source  string    `gorm:"column:source;type:varchar(30);comment:来源;not null;" json:"source"`                              // 来源
	AddTime time.Time `gorm:"column:add_time;type:datetime;comment:添加时间;not null;default:CURRENT_TIMESTAMP;" json:"add_time"` // 添加时间
	Score   uint32    `gorm:"column:score;type:int UNSIGNED;comment:添加的分数;not null;" json:"score"`                            // 添加的分数
}

func (t *UserLevelAddText) Add() error {
	return global.DB.Debug().Table("user_level_add_text").Create(&t).Error
}
