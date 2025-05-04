package model

import (
	"common/global"
	"time"
)

// 商户列表
type Merchant struct {
	MerchantId       int64     `gorm:"column:merchant_id;type:int;primaryKey;not null;" json:"merchant_id"`
	MerchantAccount  string    `gorm:"column:merchant_account;type:varchar(20);not null;" json:"merchant_account"`
	MerchantPassword string    `gorm:"column:merchant_password;type:varchar(32);not null;" json:"merchant_password"`
	ContactPhone     string    `gorm:"column:contact_phone;type:varchar(20);default:NULL;" json:"contact_phone"`
	Email            string    `gorm:"column:email;type:varchar(100);default:NULL;" json:"email"`
	RegistrationDate time.Time `gorm:"column:registration_date;type:timestamp;default:CURRENT_TIMESTAMP;" json:"registration_date"`
	Status           string    `gorm:"column:status;type:enum('active', 'inactive');default:active;" json:"status"`
}

func (m *Merchant) Create() error {
	return global.DB.Table("merchant").Create(&m).Error
}

func (m *Merchant) GetMerchantByPhone(phone string) (result *Merchant, err error) {
	err = global.DB.Table("merchant").Where("contact_phone = ?", phone).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}

func (m *Merchant) GetMerchantByEmail(email string) (result *Merchant, err error) {
	err = global.DB.Table("merchant").Where("email = ?", email).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}

func (m *Merchant) GetMerchantByAccount(account string) (result *Merchant, err error) {
	err = global.DB.Table("merchant").Where("merchant_account = ?", account).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
