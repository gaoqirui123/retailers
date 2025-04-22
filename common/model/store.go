package model

type Store struct {
	Id              uint32 `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Name            string `gorm:"column:name;type:varchar(100);comment:门店名称;not null;" json:"name"`                         // 门店名称
	Introduction    string `gorm:"column:introduction;type:varchar(1000);comment:简介;not null;" json:"introduction"`          // 简介
	Phone           string `gorm:"column:phone;type:char(25);comment:手机号码;not null;" json:"phone"`                           // 手机号码
	Address         string `gorm:"column:address;type:varchar(255);comment:省市区;not null;" json:"address"`                    // 省市区
	DetailedAddress string `gorm:"column:detailed_address;type:varchar(255);comment:详细地址;not null;" json:"detailed_address"` // 详细地址
	Image           string `gorm:"column:image;type:varchar(255);comment:门店logo;not null;" json:"image"`                     // 门店logo
	Latitude        string `gorm:"column:latitude;type:char(25);comment:纬度;not null;" json:"latitude"`                       // 纬度
	Longitude       string `gorm:"column:longitude;type:char(25);comment:经度;not null;" json:"longitude"`                     // 经度
	ValidTime       string `gorm:"column:valid_time;type:varchar(100);comment:核销有效日期;not null;" json:"valid_time"`           // 核销有效日期
	DayTime         string `gorm:"column:day_time;type:varchar(100);comment:每日营业开关时间;not null;" json:"day_time"`             // 每日营业开关时间
	AddTime         int32  `gorm:"column:add_time;type:int;comment:添加时间;not null;default:0;" json:"add_time"`                // 添加时间
	IsShow          int8   `gorm:"column:is_show;type:tinyint(1);comment:是否显示;not null;default:0;" json:"is_show"`           // 是否显示
	IsDel           int8   `gorm:"column:is_del;type:tinyint(1);comment:是否删除;not null;default:0;" json:"is_del"`             // 是否删除
}
