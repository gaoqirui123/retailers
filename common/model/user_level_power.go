package model

import (
	"common/global"
	"time"
)

// 会员权益表
type UserLevelPower struct {
	Id      uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Name    string    `gorm:"column:name;type:varchar(30);comment:权益名称;not null;" json:"name"`                                                        // 权益名称
	Grade   uint32    `gorm:"column:grade;type:int UNSIGNED;comment:会员等级：1-普通会员，2-青铜会员，3-黄铜会员，4-白银会员，5-黄金会员，6-钻石会员;not null;default:1;" json:"grade"` // 会员等级：1-普通会员，2-青铜会员，3-黄铜会员，4-白银会员，5-黄金会员，6-钻石会员
	Explain string    `gorm:"column:explain;type:text;comment:说明;not null;" json:"explain"`                                                           // 说明
	AddTime time.Time `gorm:"column:add_time;type:datetime;comment:添加时间;not null;default:CURRENT_TIMESTAMP;" json:"add_time"`                         // 添加时间
}

func (p *UserLevelPower) FindUserLevelPower() (result []UserLevelPower, err error) {
	err = global.DB.Debug().Table("user_level_power").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
