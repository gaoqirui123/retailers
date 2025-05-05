package model

import "common/global"

// 别人帮助用户砍价表记录
type BargainUserHelp struct {
	Id            uint32  `gorm:"column:id;type:int UNSIGNED;comment:砍价用户帮助表ID;primaryKey;not null;" json:"id"`                           // 砍价用户帮助表ID
	Uid           uint32  `gorm:"column:uid;type:int UNSIGNED;comment:帮助的用户id;default:NULL;" json:"uid"`                                  // 帮助者用户id
	BargainId     uint32  `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价商品ID;default:NULL;" json:"bargain_id"`                     // 砍价商品ID
	BargainUserId uint32  `gorm:"column:bargain_user_id;type:int UNSIGNED;comment:用户参与砍价表id;default:NULL;" json:"bargain_user_id"`        // 用户参与砍价表id
	Price         float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:帮助砍价多少金额;default:NULL;" json:"price"`                   // 帮助砍价多少金额
	AddTime       uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;default:NULL;" json:"add_time"`                           // 添加时间
	IsSuccess     uint8   `gorm:"column:is_success;type:tinyint UNSIGNED;comment:此次帮助砍价是否成功;default:0;" json:"is_success"`                // 此次帮助砍价是否成功:0:是,1:否
	CurrentPrice  float64 `gorm:"column:current_price;type:decimal(8, 2) UNSIGNED;comment:砍价时商品的当前价格;default:NULL;" json:"current_price"` // 砍价时商品的当前价格
}

func (BargainUserHelp) TableName() string {
	return "bargain_user_help"
}

// 创建用户帮助别人砍价表
func (m *BargainUserHelp) BargainUserHelpCreate() error {
	return global.DB.Create(&m).Error
}

// 用户帮助别人砍价表详情
func (m *BargainUserHelp) BargainUserHelpShow(Id uint32) error {
	return global.DB.Where("id = ?", Id).Find(&m).Error
}

// 砍价商品表列表
func (m *BargainUserHelp) BargainUserHelpList() (b []*BargainUserHelp, err error) {
	err = global.DB.Find(&b).Error
	if err != nil {
		return nil, err
	}
	return b, err
}
