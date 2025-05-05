package model

import (
	"common/global"
	"time"
)

// 用户提现表
type UserExtract struct {
	Id           int64     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Uid          int64     `gorm:"column:uid;type:int UNSIGNED;comment:用户id;default:NULL;" json:"uid"`                                          // 用户id
	RealName     string    `gorm:"column:real_name;type:varchar(64);comment:名称;default:NULL;" json:"real_name"`                                 // 名称
	ExtractType  string    `gorm:"column:extract_type;type:varchar(32);comment:bank = 银行卡 alipay = 支付宝wx=微信;default:bank;" json:"extract_type"` // bank = 银行卡 alipay = 支付宝wx=微信
	BankCode     string    `gorm:"column:bank_code;type:varchar(32);comment:银行卡;default:0;" json:"bank_code"`                                   // 银行卡
	BankAddress  string    `gorm:"column:bank_address;type:varchar(256);comment:开户地址;" json:"bank_address"`                                     // 开户地址
	AlipayCode   string    `gorm:"column:alipay_code;type:varchar(64);comment:支付宝账号;" json:"alipay_code"`                                       // 支付宝账号
	ExtractPrice float64   `gorm:"column:extract_price;type:decimal(8, 2) UNSIGNED;comment:提现金额;default:0.00;" json:"extract_price"`            // 提现金额
	FailMsg      string    `gorm:"column:fail_msg;type:varchar(128);comment:无效原因;default:NULL;" json:"fail_msg"`                                // 无效原因
	AddTime      time.Time `gorm:"column:add_time;type:datetime(3);comment:添加时间;not null;default:CURRENT_TIMESTAMP(3);" json:"add_time"`        // 添加时间
	Status       int64     `gorm:"column:status;type:tinyint;comment:-1 未通过 0 审核中 1 已提现;default:0;" json:"status"`                              // -1 未通过 0 审核中 1 已提现
	Wechat       string    `gorm:"column:wechat;type:varchar(15);comment:微信号;default:NULL;" json:"wechat"`                                      // 微信号
}

func (ue *UserExtract) Table() string {
	return "user_extract"
}

func (ue *UserExtract) CreateUserExtract() bool {
	err := global.DB.Debug().Table("user_extract").Create(&ue).Error
	if err != nil {
		return false
	}
	return true
}
