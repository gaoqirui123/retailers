package model

import (
	"common/global"
	"time"
)

// 用户地址表
type UserAddress struct {
	Id        int64     `gorm:"column:id;type:mediumint UNSIGNED;comment:用户地址id;primaryKey;not null;" json:"id"`                              // 用户地址id
	Uid       int64     `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                                                 // 用户id
	RealName  string    `gorm:"column:real_name;type:varchar(32);comment:收货人姓名;not null;" json:"real_name"`                                  // 收货人姓名
	Phone     string    `gorm:"column:phone;type:varchar(16);comment:收货人电话;not null;" json:"phone"`                                          // 收货人电话
	Province  string    `gorm:"column:province;type:varchar(64);comment:收货人所在省;not null;" json:"province"`                                  // 收货人所在省
	City      string    `gorm:"column:city;type:varchar(64);comment:收货人所在市;not null;" json:"city"`                                          // 收货人所在市
	CityId    int64     `gorm:"column:city_id;type:int;comment:城市id;default:NULL;" json:"city_id"`                                              // 城市id
	District  string    `gorm:"column:district;type:varchar(64);comment:收货人所在区;not null;" json:"district"`                                  // 收货人所在区
	Detail    string    `gorm:"column:detail;type:varchar(256);comment:收货人详细地址;not null;" json:"detail"`                                   // 收货人详细地址
	PostCode  int64     `gorm:"column:post_code;type:int UNSIGNED;comment:邮编;default:NULL;" json:"post_code"`                                   // 邮编
	Longitude string    `gorm:"column:longitude;type:varchar(16);comment:经度;default:0;" json:"longitude"`                                       // 经度
	Latitude  string    `gorm:"column:latitude;type:varchar(16);comment:纬度;default:0;" json:"latitude"`                                         // 纬度
	IsDefault int64     `gorm:"column:is_default;type:tinyint UNSIGNED;comment:是否默认：1-默认，2-不默认的;not null;default:1;" json:"is_default"` //是否默认：1-默认，2-不默认的
	IsDel     int64     `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;not null;default:0;" json:"is_del"`                           // 是否删除
	AddTime   time.Time `gorm:"column:add_time;type:datetime;comment:添加时间;not null;default:CURRENT_TIMESTAMP;" json:"add_time"`               // 添加时间
}

func (a *UserAddress) Created() error {
	return global.DB.Debug().Table("user_address").Create(&a).Error
}

func (a *UserAddress) FindId(userId int64) (result UserAddress, err error) {
	err = global.DB.Debug().Table("user_address").Where("uid=?", userId).Find(&result).Error
	if err != nil {
		return UserAddress{}, err
	}
	return result, nil
}

func (a *UserAddress) UpdatedAddress(userAddressId int64) error {
	return global.DB.Debug().Table("user_address").Where("id=?", userAddressId).Updates(&a).Error
}

func (a *UserAddress) FindIds(userId int64) (result []UserAddress, err error) {
	err = global.DB.Debug().Table("user_address").Where("uid=?", userId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *UserAddress) FindDefault(uid int64, userAddressId int64) (bool, error) {
	err := global.DB.Debug().Table("user_address").Where("uid=? and id=? and is_default=2", uid, userAddressId).Find(&a).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *UserAddress) UpdatedAddressDefault(userId int64) error {
	return global.DB.Debug().Table("user_address").Where("uid=?", userId).Update("is_default", 2).Error
}
