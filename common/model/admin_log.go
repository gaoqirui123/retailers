package model

import "common/global"

// EbSystemLog 管理员操作记录表
type EbSystemLog struct {
	Id         int    `gorm:"column:id;type:int UNSIGNED;comment:管理员操作记录ID;primaryKey;not null;" json:"id"`        // 管理员操作记录ID
	AdminId    int    `gorm:"column:admin_id;type:int UNSIGNED;comment:管理员id;not null;default:0;" json:"admin_id"` // 管理员id
	AdminName  string `gorm:"column:admin_name;type:varchar(64);comment:管理员姓名;not null;" json:"admin_name"`        // 管理员姓名
	Path       string `gorm:"column:path;type:varchar(128);comment:链接;not null;" json:"path"`                      // 链接
	Page       string `gorm:"column:page;type:varchar(64);comment:行为;not null;" json:"page"`                       // 行为
	Method     string `gorm:"column:method;type:varchar(12);comment:访问类型;not null;" json:"method"`                 // 访问类型
	Ip         string `gorm:"column:ip;type:varchar(16);comment:登录IP;not null;" json:"ip"`                         // 登录IP
	Type       string `gorm:"column:type;type:varchar(32);comment:类型;not null;" json:"type"`                       // 类型
	AddTime    int    `gorm:"column:add_time;type:int UNSIGNED;comment:操作时间;default:0;" json:"add_time"`           // 操作时间
	MerchantId int    `gorm:"column:merchant_id;type:int UNSIGNED;comment:商户id;default:0;" json:"merchant_id"`     // 商户id
}

func (l *EbSystemLog) Create() error {
	return global.DB.Debug().Table("eb_system_log").Create(&l).Error
}
