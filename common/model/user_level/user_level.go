package user_level

import (
	"common/global"
	"time"
)

// 会员表
type UserLevel struct {
	Id           int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	MerId        int32     `gorm:"column:mer_id;type:int;comment:商户id;not null;default:0;" json:"mer_id"`                          // 商户id
	Name         string    `gorm:"column:name;type:varchar(255);comment:会员名称;not null;" json:"name"`                               // 会员名称
	MemberPoints string    `gorm:"column:member_points;type:varchar(100);comment:会员积分;not null;default:0;" json:"member_points"`   // 会员积分
	IsShow       int8      `gorm:"column:is_show;type:tinyint(1);comment:是否显示 1=显示,0=隐藏;not null;default:0;" json:"is_show"`       // 是否显示 1=显示,0=隐藏
	Grade        int32     `gorm:"column:grade;type:int;comment:会员等级;not null;default:0;" json:"grade"`                            // 会员等级
	Image        string    `gorm:"column:image;type:varchar(255);comment:会员卡背景;not null;" json:"image"`                            // 会员卡背景
	Icon         string    `gorm:"column:icon;type:varchar(255);comment:会员图标;not null;" json:"icon"`                               // 会员图标
	Explain      string    `gorm:"column:explain;type:text;comment:说明;not null;" json:"explain"`                                   // 说明
	AddTime      time.Time `gorm:"column:add_time;type:datetime;comment:添加时间;not null;default:CURRENT_TIMESTAMP;" json:"add_time"` // 添加时间
}

func (l *UserLevel) FindUsersLevel() (result []UserLevel, err error) {
	err = global.DB.Debug().Table("user_level").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
