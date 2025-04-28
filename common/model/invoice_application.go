package model

import (
	"common/global"
	"time"
)

type InvoiceApplication struct {
	ApplicationId                int64     `gorm:"column:application_id;type:int;primaryKey;not null;" json:"application_id"`
	UserId                       int64     `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                                                             // 用户id
	OrderId                      int64     `gorm:"column:order_id;type:int;comment:订单id;not null;" json:"order_id"`                                                           // 订单id
	InvoiceType                  string    `gorm:"column:invoice_type;type:varchar(50);comment:发票类型：普通发票、增值税专用发票;not null;" json:"invoice_type"`                              // 发票类型：普通发票、增值税专用发票
	InvoiceTitle                 string    `gorm:"column:invoice_title;type:varchar(255);comment:发票抬头;not null;" json:"invoice_title"`                                        // 发票抬头
	TaxpayerIdentificationNumber string    `gorm:"column:taxpayer_identification_number;type:varchar(20);comment:纳税人识别号;default:NULL;" json:"taxpayer_identification_number"` // 纳税人识别号
	InvoiceAmount                float64   `gorm:"column:invoice_amount;type:decimal(10, 2);comment:发票金额;not null;" json:"invoice_amount"`                                    // 发票金额
	ApplicationTime              time.Time `gorm:"column:application_time;type:timestamp;comment:申请时间;default:CURRENT_TIMESTAMP;" json:"application_time"`                    // 申请时间
	ApplicationStatus            int64     `gorm:"column:application_status;type:tinyint;comment:申请状态：1待审核、2已通过、3已拒绝;default:1;" json:"application_status"`                   // 申请状态：1待审核、2已通过、3已拒绝
	ReviewTime                   time.Time `gorm:"column:review_time;type:timestamp;comment:审核时间;default:NULL;" json:"review_time"`                                           // 审核时间
	ReviewRemark                 string    `gorm:"column:review_remark;type:text;comment:审核备注;" json:"review_remark"`                                                         // 审核备注
	Email                        string    `gorm:"column:email;type:varchar(20);comment:邮箱;not null;" json:"email"`                                                           // 邮箱
	Address                      string    `gorm:"column:address;type:varchar(255);comment:寄件地址;not null;" json:"address"`                                                    // 寄件地址
	Phone                        string    `gorm:"column:phone;type:char(11);comment:联系人电话;not null;" json:"phone"`                                                           // 联系人电话
	Type                         string    `gorm:"column:type;type:varchar(20);comment:发票材质：纸质、电子;not null;" json:"type"`                                                     // 发票材质：纸质、电子
	MerId                        int64     `gorm:"column:mer_id;type:int;comment:商户id;not null;" json:"mer_id"`                                                               // 商户id
}

func (a *InvoiceApplication) UserApplication() error {
	return global.DB.Debug().Table("invoice_application").Create(&a).Error
}
