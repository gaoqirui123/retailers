package model

// 会员等级记录表
type EbUserLevel struct {
	Id        int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Uid       int32  `gorm:"column:uid;type:int;comment:用户uid;not null;default:0;" json:"uid"`                        // 用户uid
	LevelId   int32  `gorm:"column:level_id;type:int;comment:等级vip;not null;default:0;" json:"level_id"`              // 等级vip
	Grade     int32  `gorm:"column:grade;type:int;comment:会员等级;not null;default:0;" json:"grade"`                     // 会员等级
	ValidTime int32  `gorm:"column:valid_time;type:int;comment:过期时间;not null;default:0;" json:"valid_time"`           // 过期时间
	IsForever int8   `gorm:"column:is_forever;type:tinyint(1);comment:是否永久;not null;default:0;" json:"is_forever"`    // 是否永久
	MerId     int32  `gorm:"column:mer_id;type:int;comment:商户id;not null;default:0;" json:"mer_id"`                   // 商户id
	Status    int8   `gorm:"column:status;type:tinyint(1);comment:0:禁止,1:正常;not null;default:0;" json:"status"`       // 0:禁止,1:正常
	Mark      string `gorm:"column:mark;type:varchar(255);comment:备注;not null;" json:"mark"`                          // 备注
	Remind    int8   `gorm:"column:remind;type:tinyint(1);comment:是否已通知;not null;default:0;" json:"remind"`           // 是否已通知
	IsDel     int8   `gorm:"column:is_del;type:tinyint(1);comment:是否删除,0=未删除,1=删除;not null;default:0;" json:"is_del"` // 是否删除,0=未删除,1=删除
	AddTime   int32  `gorm:"column:add_time;type:int;comment:添加时间;not null;default:0;" json:"add_time"`               // 添加时间
	Discount  int32  `gorm:"column:discount;type:int;comment:享受折扣;not null;default:0;" json:"discount"`               // 享受折扣
}
