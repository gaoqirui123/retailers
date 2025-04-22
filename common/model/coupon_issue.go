package model

type CouponIssue struct {
	Id              uint32  `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Cid             int32   `gorm:"column:cid;type:int;comment:优惠券ID;default:NULL;" json:"cid"`                                                         // 优惠券ID
	StartTime       int32   `gorm:"column:start_time;type:int;comment:优惠券领取开启时间;default:NULL;" json:"start_time"`                                       // 优惠券领取开启时间
	EndTime         int32   `gorm:"column:end_time;type:int;comment:优惠券领取结束时间;default:NULL;" json:"end_time"`                                           // 优惠券领取结束时间
	TotalCount      int32   `gorm:"column:total_count;type:int;comment:优惠券领取数量;default:NULL;" json:"total_count"`                                       // 优惠券领取数量
	RemainCount     int32   `gorm:"column:remain_count;type:int;comment:优惠券剩余领取数量;default:NULL;" json:"remain_count"`                                   // 优惠券剩余领取数量
	IsPermanent     int8    `gorm:"column:is_permanent;type:tinyint(1);comment:是否无限张数;not null;default:0;" json:"is_permanent"`                         // 是否无限张数
	Status          int8    `gorm:"column:status;type:tinyint(1);comment:1 正常 0 未开启 -1 已无效;not null;default:1;" json:"status"`                          // 1 正常 0 未开启 -1 已无效
	IsGiveSubscribe int8    `gorm:"column:is_give_subscribe;type:tinyint(1);comment:是否首次关注赠送 0-否(默认) 1-是;not null;default:0;" json:"is_give_subscribe"` // 是否首次关注赠送 0-否(默认) 1-是
	IsFullGive      int8    `gorm:"column:is_full_give;type:tinyint(1);comment:是否满赠0-否(默认) 1-是;not null;default:0;" json:"is_full_give"`                // 是否满赠0-否(默认) 1-是
	FullReduction   float64 `gorm:"column:full_reduction;type:decimal(8, 2);comment:消费满多少赠送优惠券;not null;default:0.00;" json:"full_reduction"`           // 消费满多少赠送优惠券
	IsDel           uint8   `gorm:"column:is_del;type:tinyint UNSIGNED;not null;default:0;" json:"is_del"`
	AddTime         int32   `gorm:"column:add_time;type:int;comment:优惠券添加时间;default:NULL;" json:"add_time"` // 优惠券添加时间
}
