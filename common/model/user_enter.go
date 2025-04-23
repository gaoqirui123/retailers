package model

import "common/global"

// UserEnter 商户申请表
type UserEnter struct {
	Id           int    `gorm:"column:id;type:int UNSIGNED;comment:商户申请ID;primaryKey;not null;" json:"id"`          // 商户申请ID
	Uid          int    `gorm:"column:uid;type:int UNSIGNED;comment:用户ID;not null;default:0;" json:"uid"`           // 用户ID
	Province     string `gorm:"column:province;type:varchar(32);comment:商户所在省;not null;" json:"province"`           // 商户所在省
	City         string `gorm:"column:city;type:varchar(32);comment:商户所在市;not null;" json:"city"`                   // 商户所在市
	District     string `gorm:"column:district;type:varchar(32);comment:商户所在区;not null;" json:"district"`           // 商户所在区
	Address      string `gorm:"column:address;type:varchar(256);comment:商户详细地址;not null;" json:"address"`           // 商户详细地址
	MerchantName string `gorm:"column:merchant_name;type:varchar(256);comment:商户名称;not null;" json:"merchant_name"` // 商户名称
	LinkUser     string `gorm:"column:link_user;type:varchar(32);not null;" json:"link_user"`
	LinkTel      string `gorm:"column:link_tel;type:varchar(16);comment:商户电话;not null;" json:"link_tel"`                      // 商户电话
	Charter      string `gorm:"column:charter;type:varchar(512);comment:商户证书;not null;" json:"charter"`                       // 商户证书
	AddTime      int    `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;not null;default:0;" json:"add_time"`           // 添加时间
	ApplyTime    int    `gorm:"column:apply_time;type:int UNSIGNED;comment:审核时间;not null;default:0;" json:"apply_time"`       // 审核时间
	SuccessTime  int    `gorm:"column:success_time;type:int;comment:通过时间;not null;" json:"success_time"`                      // 通过时间
	FailMessage  string `gorm:"column:fail_message;type:varchar(256);comment:未通过原因;not null;" json:"fail_message"`            // 未通过原因
	FailTime     int    `gorm:"column:fail_time;type:int UNSIGNED;comment:未通过时间;not null;default:0;" json:"fail_time"`        // 未通过时间
	Status       int    `gorm:"column:status;type:tinyint(1);comment:-1 审核未通过 0未审核 1审核通过;not null;default:0;" json:"status"`  // -1 审核未通过 0未审核 1审核通过
	IsLock       int    `gorm:"column:is_lock;type:tinyint UNSIGNED;comment:0 = 开启 1= 关闭;not null;default:0;" json:"is_lock"` // 0 = 开启 1= 关闭
	IsDel        int    `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;not null;default:0;" json:"is_del"`           // 是否删除
}

func (ue *UserEnter) TableName() string {
	return "user_enter"
}

func (ue *UserEnter) Add() error {
	return global.DB.Create(&ue).Error
}

func (ue *UserEnter) UpdateStatus(Id int64, status int64) error {
	return global.DB.Table("user_enter").Where("id = ?", Id).Update("status", status).Error
}

func (ue *UserEnter) GetStatusById(id int64) (result *UserEnter, err error) {
	err = global.DB.Table("user_enter").Where("id = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
