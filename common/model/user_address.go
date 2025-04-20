package model

type UserAddress struct {
	Id        uint32 `gorm:"column:id;type:mediumint UNSIGNED;comment:用户地址id;primaryKey;not null;" json:"id"`            // 用户地址id
	Uid       uint32 `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                             // 用户id
	RealName  string `gorm:"column:real_name;type:varchar(32);comment:收货人姓名;not null;" json:"real_name"`                 // 收货人姓名
	Phone     string `gorm:"column:phone;type:varchar(16);comment:收货人电话;not null;" json:"phone"`                         // 收货人电话
	Province  string `gorm:"column:province;type:varchar(64);comment:收货人所在省;not null;" json:"province"`                  // 收货人所在省
	City      string `gorm:"column:city;type:varchar(64);comment:收货人所在市;not null;" json:"city"`                          // 收货人所在市
	CityId    int32  `gorm:"column:city_id;type:int;comment:城市id;not null;default:0;" json:"city_id"`                    // 城市id
	District  string `gorm:"column:district;type:varchar(64);comment:收货人所在区;not null;" json:"district"`                  // 收货人所在区
	Detail    string `gorm:"column:detail;type:varchar(256);comment:收货人详细地址;not null;" json:"detail"`                    // 收货人详细地址
	PostCode  uint32 `gorm:"column:post_code;type:int UNSIGNED;comment:邮编;not null;" json:"post_code"`                   // 邮编
	Longitude string `gorm:"column:longitude;type:varchar(16);comment:经度;not null;default:0;" json:"longitude"`          // 经度
	Latitude  string `gorm:"column:latitude;type:varchar(16);comment:纬度;not null;default:0;" json:"latitude"`            // 纬度
	IsDefault uint8  `gorm:"column:is_default;type:tinyint UNSIGNED;comment:是否默认;not null;default:0;" json:"is_default"` // 是否默认
	IsDel     uint8  `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;not null;default:0;" json:"is_del"`         // 是否删除
	AddTime   uint32 `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;not null;default:0;" json:"add_time"`         // 添加时间
}
