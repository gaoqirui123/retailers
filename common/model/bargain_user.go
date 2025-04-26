package model

// 用户参与砍价商品记录表（只要砍过都算，不管是否砍成功）
type BargainUser struct {
	Id              uint32  `gorm:"column:id;type:int UNSIGNED;comment:用户参与砍价表ID;primaryKey;not null;" json:"id"`                               // 用户参与砍价表ID
	Uid             uint32  `gorm:"column:uid;type:int UNSIGNED;comment:用户ID;default:NULL;" json:"uid"`                                         // 用户ID
	BargainId       uint32  `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价商品id;default:NULL;" json:"bargain_id"`                         // 砍价商品id
	BargainPriceMin float64 `gorm:"column:bargain_price_min;type:decimal(8, 2) UNSIGNED;comment:砍价的最低价;default:NULL;" json:"bargain_price_min"` // 砍价的最低价
	BargainPrice    float64 `gorm:"column:bargain_price;type:decimal(8, 2);comment:砍价金额;default:NULL;" json:"bargain_price"`                    // 砍价金额
	Price           float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:砍掉的价格;default:NULL;" json:"price"`                          // 砍掉的价格
	Status          uint8   `gorm:"column:status;type:tinyint UNSIGNED;comment:状态 1参与中 2 活动结束参与失败 3活动结束参与成功;not null;default:0;" json:"status"` // 状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	AddTime         uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:参与时间;default:NULL;" json:"add_time"`                               // 参与时间
	IsDel           int8    `gorm:"column:is_del;type:tinyint(1);comment:是否取消;not null;default:0;" json:"is_del"`                               // 是否取消
}

func (BargainUser) TableName() string {
	return "bargain_user"
}
