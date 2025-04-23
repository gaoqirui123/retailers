package model

import "common/global"

// EbSystemAdmin 后台管理员表
type EbSystemAdmin struct {
	Id         int    `gorm:"column:id;type:smallint UNSIGNED;comment:后台管理员表ID;primaryKey;not null;" json:"id"`             // 后台管理员表ID
	Account    string `gorm:"column:account;type:varchar(32);comment:后台管理员账号;not null;" json:"account"`                     // 后台管理员账号
	Pwd        string `gorm:"column:pwd;type:char(32);comment:后台管理员密码;not null;" json:"pwd"`                                // 后台管理员密码
	RealName   string `gorm:"column:real_name;type:varchar(16);comment:后台管理员姓名;not null;" json:"real_name"`                 // 后台管理员姓名
	Roles      string `gorm:"column:roles;type:varchar(128);comment:后台管理员权限(menus_id);not null;" json:"roles"`              // 后台管理员权限(menus_id)
	LastIp     string `gorm:"column:last_ip;type:varchar(16);comment:后台管理员最后一次登录ip;default:NULL;" json:"last_ip"`           // 后台管理员最后一次登录ip
	LastTime   uint32 `gorm:"column:last_time;type:int UNSIGNED;comment:后台管理员最后一次登录时间;default:NULL;" json:"last_time"`      // 后台管理员最后一次登录时间
	AddTime    int    `gorm:"column:add_time;type:int UNSIGNED;comment:后台管理员添加时间;not null;default:0;" json:"add_time"`      // 后台管理员添加时间
	LoginCount int    `gorm:"column:login_count;type:int UNSIGNED;comment:登录次数;not null;default:0;" json:"login_count"`     // 登录次数
	Level      int    `gorm:"column:level;type:tinyint UNSIGNED;comment:后台管理员级别;not null;default:1;" json:"level"`          // 后台管理员级别
	Status     int    `gorm:"column:status;type:tinyint UNSIGNED;comment:后台管理员状态 1有效0无效;not null;default:1;" json:"status"` // 后台管理员状态 1有效0无效
	IsDel      int    `gorm:"column:is_del;type:tinyint UNSIGNED;not null;default:0;" json:"is_del"`
}

func (a *EbSystemAdmin) GetAdminByAcoount(account string) (result *EbSystemAdmin, err error) {
	err = global.DB.Table("eb_system_admin").Where("account = ?", account).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}

func (a *EbSystemAdmin) GetAdminById(id int64) (result *EbSystemAdmin, err error) {
	err = global.DB.Table("eb_system_admin").Where("id = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
