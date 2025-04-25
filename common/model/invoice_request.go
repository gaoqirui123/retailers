package model

// InvoiceRequest 发票申请表
type InvoiceRequest struct {
	Id      int64   `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	UserId  int64   `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                         // 用户id
	Amount  float64 `gorm:"column:amount;type:decimal(10, 2);comment:价格;not null;" json:"amount"`                  // 价格
	OrderSn string  `gorm:"column:order_sn;type:varchar(255);comment:订单号;not null;" json:"order_sn"`               // 订单号
	Status  int64   `gorm:"column:status;type:int;comment:申请状态 1待审批,2 已批准,3已拒绝;not null;default:1;" json:"status"` // 申请状态 1待审批,2 已批准,3已拒绝
	Reason  string  `gorm:"column:reason;type:varchar(255);comment:原因;default:NULL;" json:"reason"`                // 原因
	Type    string  `gorm:"column:type;type:varchar(50);comment:类型;not null;" json:"type"`                         // 类型
}
