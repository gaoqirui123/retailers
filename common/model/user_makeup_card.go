package model

import (
	"common/global"
	"gorm.io/gorm"
	"time"
)

//type UserMakeup struct {
//	Id         int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
//	UserId     int64     `gorm:"column:user_id;type:bigint;comment:用户ID;not null;" json:"user_id"`                  // 用户ID
//	Cardcount  int64     `gorm:"column:cardCount;type:bigint;comment:补签卡;default:0;" json:"cardCount"`              // 补签卡
//	UpdateTime time.Time `gorm:"column:update_time;type:datetime(3);comment:修改时间;default:NULL;" json:"update_time"` // 修改时间
//}

type UserMakeup struct {
	Id         int64     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	UserId     int64     `gorm:"column:user_id;type:int UNSIGNED;comment:用户ID;not null;" json:"user_id"`            // 用户ID
	Cardcount  int64     `gorm:"column:cardCount;type:int UNSIGNED;comment:补签卡;not null;" json:"cardCount"`         // 补签卡
	UpdateTime time.Time `gorm:"column:update_time;type:datetime(3);comment:修改时间;default:NULL;" json:"update_time"` // 修改时间
}

func (um *UserMakeup) GetUserMakeupCard(userId int64) error {
	return global.DB.Debug().Table("user_makeup").Where("user_id = ?  AND cardCount > 0", userId).Limit(1).Find(&um).Error
}

func (um *UserMakeup) UpdateUserMakeupCard(userId int64) error {
	return global.DB.Debug().Table("user_makeup").Model(&UserMakeup{}).Where("user_id = ? AND cardCount > 0", userId).Update("cardCount", gorm.Expr("cardCount - 1")).Update("update_time", time.Now()).Error
}
