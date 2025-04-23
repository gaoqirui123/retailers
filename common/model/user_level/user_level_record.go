package user_level

import (
	"common/global"
	"time"
)

// 用户会员记录表
type UserLevelRecord struct {
	Id        int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Uid       int32     `gorm:"column:uid;type:int;comment:用户uid;not null;default:0;" json:"uid"`                                   // 用户uid
	Grade     int32     `gorm:"column:grade;type:int;comment:会员等级;not null;default:0;" json:"grade"`                                // 会员等级
	ValidTime time.Time `gorm:"column:valid_time;type:datetime;comment:过期时间;not null;default:CURRENT_TIMESTAMP;" json:"valid_time"` // 过期时间
	MerId     int32     `gorm:"column:mer_id;type:int;comment:商户id;not null;default:0;" json:"mer_id"`                              // 商户id
	Status    int8      `gorm:"column:status;type:tinyint(1);comment:0:禁止,1:正常;not null;default:0;" json:"status"`                  // 0:禁止,1:正常
	Mark      string    `gorm:"column:mark;type:varchar(255);comment:备注;not null;" json:"mark"`                                     // 备注
	Remind    int8      `gorm:"column:remind;type:tinyint(1);comment:是否已通知;not null;default:0;" json:"remind"`                      // 是否已通知
	AddTime   time.Time `gorm:"column:add_time;type:datetime;comment:添加时间;not null;default:CURRENT_TIMESTAMP;" json:"add_time"`     // 添加时间
	Discount  int32     `gorm:"column:discount;type:int;comment:享受折扣;not null;default:0;" json:"discount"`                          // 享受折扣
}

func (r *UserLevelRecord) FindRecords(uid int) (result UserLevelRecord, err error) {
	err = global.DB.Debug().Table("user_level_record").Where("uid=?", uid).Find(&result).Error
	if err != nil {
		return UserLevelRecord{}, err
	}
	return result, nil
}
