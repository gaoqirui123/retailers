package user_level

import (
	"common/global"
	"time"
)

// 用户使用权益表
type UserLevelUsePower struct {
	Id      uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Uid     uint32    `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                                 // 用户id
	Qid     uint32    `gorm:"column:qid;type:int UNSIGNED;comment:权益id;not null;" json:"qid"`                                 // 权益id
	AddTime time.Time `gorm:"column:add_time;type:datetime;comment:添加时间;not null;default:CURRENT_TIMESTAMP;" json:"add_time"` // 添加时间
}

func (p *UserLevelUsePower) AddUserPower() error {
	return global.DB.Debug().Table("user_level_use_power").Create(&p).Error
}

func (p *UserLevelUsePower) Finds() (result []UserLevelUsePower, err error) {
	err = global.DB.Debug().Table("user_level_use_power").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
