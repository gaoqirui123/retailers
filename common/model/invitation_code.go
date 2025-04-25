package model

import (
	"common/global"
	"time"
)

// 邀请码表
type InvitationCode struct {
	Id        uint64    `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	Uid       int64     `gorm:"column:uid;type:bigint;comment:用户id;not null;" json:"uid"`                                                 // 用户id
	Code      string    `gorm:"column:code;type:varchar(255);comment:邀请码;not null;" json:"code"`                                          // 邀请码
	Status    int64     `gorm:"column:status;type:bigint;comment:1未使用2已使用/已过期;not null;default:1;" json:"status"`                         // 1未使用2已使用
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:添加时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 添加时间
	Type      int64     `gorm:"column:type;type:bigint;comment:1邀请码2链接3二维码;not null;default:1;" json:"type"`                              // 1邀请码2链接3二维码
	IsDel     int64     `gorm:"column:is_del;type:bigint;comment:1未删除2已删除;not null;default:1;" json:"is_del"`                             // 1未删除2已删除
}

func (i *InvitationCode) Table() string {
	return "invitation_code"
}

func (i *InvitationCode) Create() bool {
	err := global.DB.Debug().Table("invitation_code").Create(&i).Error
	if err != nil {
		return false
	}

	return true
}

// 查找邀请码不存在
func (i *InvitationCode) FindCode(code string) InvitationCode {
	var in InvitationCode
	err := global.DB.Debug().Table("invitation_code").Where("code = ?", code).Find(&in).Error
	if err != nil {
		return InvitationCode{}
	}
	return in
}

// 修改邀请码状态
func (i *InvitationCode) UpdateCode(code string) bool {

	err := global.DB.Debug().Table("invitation_code").Where("code = ?", code).Update("status", 2).Error

	if err != nil {
		return false
	}
	return true
}

// 删除邀请码
func (i *InvitationCode) DeleteCode(code string) bool {

	err := global.DB.Debug().Table("invitation_code").Where("code = ?", code).Update("is_del", 2).Error

	if err != nil {
		return false
	}
	return true
}
