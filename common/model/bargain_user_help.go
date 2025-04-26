package model

// 用户帮助别人砍价表
type BargainUserHelp struct {
	Id            uint32  `gorm:"column:id;type:int UNSIGNED;comment:砍价用户帮助表ID;primaryKey;not null;" json:"id"`                    // 砍价用户帮助表ID
	Uid           uint32  `gorm:"column:uid;type:int UNSIGNED;comment:帮助的用户id;default:NULL;" json:"uid"`                           // 帮助的用户id
	BargainId     uint32  `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价商品ID;default:NULL;" json:"bargain_id"`              // 砍价商品ID
	BargainUserId uint32  `gorm:"column:bargain_user_id;type:int UNSIGNED;comment:用户参与砍价表id;default:NULL;" json:"bargain_user_id"` // 用户参与砍价表id
	Price         float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:帮助砍价多少金额;default:NULL;" json:"price"`            // 帮助砍价多少金额
	AddTime       uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;default:NULL;" json:"add_time"`                    // 添加时间
}

func (BargainUserHelp) TableName() string {
	return "bargain_user_help"
}
