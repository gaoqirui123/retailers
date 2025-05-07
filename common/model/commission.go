package model

import (
	"common/global"
	"log"
	"time"
)

// 返佣流水表
type Commission struct {
	CommissionId uint32    `gorm:"column:commission_id;type:int UNSIGNED;comment:返佣流水ID（主键）;primaryKey;not null;" json:"commission_id"`                       // 返佣流水ID（主键）
	OrderSyn     string    `gorm:"column:order_syn;type:varchar(255);comment:触发返佣的订单号;not null;" json:"order_syn"`                                            // 触发返佣的订单号
	FromUserId   uint32    `gorm:"column:from_user_id;type:int UNSIGNED;comment:消费用户ID;not null;" json:"from_user_id"`                                        // 消费用户ID
	ToUserId     uint32    `gorm:"column:to_user_id;type:int UNSIGNED;comment:获得佣金用户ID;not null;" json:"to_user_id"`                                          // 获得佣金用户ID
	Level        int8      `gorm:"column:level;type:tinyint(1);comment:佣金层级：1=一级返佣, 2=二级返佣;not null;" json:"level"`                                           // 佣金层级：1=一级返佣, 2=二级返佣
	Amount       float64   `gorm:"column:amount;type:decimal(10, 2);comment:佣金金额（精度：小数点后两位）;not null;" json:"amount"`                                         // 佣金金额（精度：小数点后两位）
	Status       string    `gorm:"column:status;type:enum('normal', 'recovered');comment:状态：normal=正常, recovered=已追回;not null;default:normal;" json:"status"` // 状态：normal=正常, recovered=已追回
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"create_time"`                      // 创建时间
}

func (n *Commission) Table() string {
	return "commission"
}

func (n *Commission) CreateCommission() bool {
	err := global.DB.Debug().Table("commission").Create(&n).Error
	if err != nil {
		return false
	}
	return true
}

// 计算用户总佣金并按佣金从高到低排序
func (n *Commission) CalculateAndRankTotalCommission() []struct {
	TotalAmount float64 `json:"total_amount"`
	Account     string  `json:"account"` // 用户名字段
	Avatar      string  `json:"avatar"`
} {
	var result []struct {
		TotalAmount float64 `json:"total_amount"`
		Account     string  `json:"account"` // 用户名字段
		Avatar      string  `json:"avatar"`
	}
	// 使用更高效的查询方式，移除 Debug 模式，并关联 user 表
	query := global.DB.Table("commission").
		Select("commission.to_user_id, SUM(commission.amount) as total_amount, user.account,user.avatar").
		Joins("JOIN user ON commission.to_user_id = user.uid").
		Where("commission.status = 'normal'").
		Group("commission.to_user_id").
		Order("total_amount DESC")
	err := query.Find(&result).Error
	if err != nil {

		log.Printf("Failed to calculate and rank total commission: %v", err)
	}
	// 打印完整结果
	//	log.Printf("Query result: %+v", result)
	return result
}

func (n *Commission) RankPrice() []Commission {
	var commission []Commission
	err := global.DB.Debug().Table("commission").Find(&commission).Error
	if err != nil {
		return commission
	}
	return commission
}
