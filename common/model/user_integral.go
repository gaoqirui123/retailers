package model

import (
	"common/global"
	"gorm.io/gorm"
	"time"
)

// 积分类型常量
const (
	IntegralTypeSignIn     = "1" // 签到
	IntegralTypeContinuous = "2" // 连续签到
	IntegralTypeWelfare    = "3" // 福利任务
	IntegralTypeDaily      = "4" // 每日任务
	IntegralTypeReplenish  = "5" // 补签
)

type UserIntegral struct {
	Id            int64     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	UserId        int64     `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                     // 用户id
	Integral      int64     `gorm:"column:integral;type:int;comment:当前积分;default:0;" json:"integral"`                  // 当前积分
	IntegralTotal int64     `gorm:"column:integral_total;type:int;comment:累计积分;default:0;" json:"integral_total"`      // 累计积分
	CreateTime    time.Time `gorm:"column:create_time;type:datetime(3);comment: 创建时间;not null;" json:"create_time"`    //  创建时间
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime(3);comment:修改时间;default:NULL;" json:"update_time"` // 修改时间
}

func (ui *UserIntegral) GetUserIntegral(userId int64) error {
	return global.DB.Debug().Table("user_integral").Where("user_id = ?", userId).Limit(1).Find(&ui).Error
}
func (ui *UserIntegral) AddUserIntegral() error {
	return global.DB.Debug().Table("user_integral").Create(&ui).Error
}

func (ui *UserIntegral) UpdateUserIntegral(userId, points int64) error {
	return global.DB.Debug().Table("user_integral").Model(&UserIntegral{}).Where("user_id = ?", userId).Updates(map[string]interface{}{"integral": gorm.Expr("integral + ?", points), "integral_total": gorm.Expr("integral_total + ?", points), "update_time": time.Now()}).Error
}
