package model

import (
	"common/global"
	"gorm.io/gorm"
	"time"
)

type UserMakeupCard struct {
	Id         int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	UserId     int64     `gorm:"column:user_id;type:bigint;comment:用户ID;not null;" json:"user_id"`                  // 用户ID
	Cardcount  int64     `gorm:"column:cardCount;type:bigint;comment:补签卡;default:0;" json:"cardCount"`              // 补签卡
	UpdateTime time.Time `gorm:"column:update_time;type:datetime(3);comment:修改时间;default:NULL;" json:"update_time"` // 修改时间
}

func (um *UserMakeupCard) GetUserMakeupCard(userId int64) error {
	return global.DB.Debug().Table("user_makeup_cart").Where("user_id = ?  AND cardCount > 0", userId).Limit(1).Find(&um).Error
}

func (um *UserMakeupCard) UpdateUserMakeupCard(userId int64) error {
	return global.DB.Debug().Table("user_makeup_cart").Model(&UserMakeupCard{}).Where("user_id = ? AND cardCount > 0", userId).Update("cardCount", gorm.Expr("cardCount - 1")).Update("update_time", time.Now()).Error
}
