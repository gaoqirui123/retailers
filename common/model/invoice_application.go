package model

import (
	"common/global"
	"common/proto/user_enter"
	"time"
)

// InvoiceApplication 发票申请表
type InvoiceApplication struct {
	ApplicationId                int64     `gorm:"column:application_id;type:int;primaryKey;not null;" json:"application_id"`
	UserId                       int64     `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                                                             // 用户id
	OrderId                      int64     `gorm:"column:order_id;type:int;comment:订单id;not null;" json:"order_id"`                                                           // 订单id
	InvoiceType                  string    `gorm:"column:invoice_type;type:varchar(50);comment:发票类型：普通发票、增值税专用发票;not null;" json:"invoice_type"`                              // 发票类型：普通发票、增值税专用发票
	InvoiceTitle                 string    `gorm:"column:invoice_title;type:varchar(255);comment:发票抬头;not null;" json:"invoice_title"`                                        // 发票抬头
	TaxpayerIdentificationNumber string    `gorm:"column:taxpayer_identification_number;type:varchar(20);comment:纳税人识别号;default:NULL;" json:"taxpayer_identification_number"` // 纳税人识别号
	InvoiceAmount                float64   `gorm:"column:invoice_amount;type:decimal(10, 2);comment:发票金额;not null;" json:"invoice_amount"`                                    // 发票金额
	ApplicationTime              time.Time `gorm:"column:application_time;type:timestamp;comment:申请时间;default:CURRENT_TIMESTAMP;" json:"application_time"`                    // 申请时间
	ApplicationStatus            int8      `gorm:"column:application_status;type:tinyint;comment:申请状态：1待审核、2已通过、3已拒绝;default:1;" json:"application_status"`                   // 申请状态：1待审核、2已通过、3已拒绝
	ReviewTime                   time.Time `gorm:"column:review_time;type:timestamp;comment:审核时间;default:NULL;" json:"review_time"`                                           // 审核时间
	ReviewRemark                 string    `gorm:"column:review_remark;type:text;comment:审核备注;" json:"review_remark"`                                                         // 审核备注
	Email                        string    `gorm:"column:email;type:varchar(20);comment:邮箱;not null;" json:"email"`                                                           // 邮箱
	Address                      string    `gorm:"column:address;type:varchar(255);comment:寄件地址;not null;" json:"address"`                                                    // 寄件地址
	Phone                        string    `gorm:"column:phone;type:char(11);comment:联系人电话;not null;" json:"phone"`                                                           // 联系人电话
	Type                         string    `gorm:"column:type;type:varchar(20);comment:发票材质：纸质、电子;not null;" json:"type"`                                                     // 发票材质：纸质、电子
	MerId                        int64     `gorm:"column:mer_id;type:int;comment:商户id;not null;" json:"mer_id"`                                                               // 商户id
	ProductName                  string    `gorm:"column:product_name;type:varchar(255);comment:商品名称;not null;" json:"product_name"`                                          // 商品名称
	ProductNumber                int64     `gorm:"column:product_number;type:int;comment:商品数量;not null;" json:"product_number"`                                               // 商品数量
	ProductPrice                 float64   `gorm:"column:product_price;type:decimal(10, 2);comment:商品单价;not null;" json:"product_price"`                                      // 商品单价
	ProductAmount                float64   `gorm:"column:product_amount;type:decimal(10, 2);comment:商品总价;not null;" json:"product_amount"`                                    // 商品总价
}

func (a *InvoiceApplication) UserApplication() error {
	return global.DB.Debug().Table("invoice_application").Create(&a).Error
}
func (a *InvoiceApplication) GetInvoiceByUeIds(ueId int64) (result []*InvoiceApplication, err error) {
	err = global.DB.Table("invoice_application").Where("mer_id = ?", ueId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
func (a *InvoiceApplication) GetInvoiceByUeId(uId int64, orderId int64) (result *InvoiceApplication, err error) {
	err = global.DB.Table("invoice_application").Where("user_id  = ?", uId).Where("order_id = ?", orderId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
func (a *InvoiceApplication) GetInvoiceByUeIdAndStatus(ueId int64, status int64) (result []*InvoiceApplication, err error) {
	err = global.DB.Table("invoice_application").Where("mer_id = ?", ueId).Where("application_status = ?", status).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}

func (a *InvoiceApplication) UpdateStatus(ueId int64, uid int64, status int64, orderId int64, now time.Time) error {
	return global.DB.Table("invoice_application").Where("mer_id = ?", ueId).Where("user_id = ?", uid).Where("order_id = ?", orderId).Update("application_status", status).Update("review_remark", now).Error
}

func (a *InvoiceApplication) UpdateStatusDis(ueId int64, uid int64, status int64, dis string, orderId int64, now time.Time) error {
	err := global.DB.Table("invoice_application").Where("mer_id = ?", ueId).Where("user_id = ?", uid).Where("order_id = ?", orderId).Update("application_status", status).Update("review_remark", now).Error
	if err != nil {
		return err
	}
	err = global.DB.Table("invoice_application").Where("mer_id = ?", ueId).Where("user_id = ?", uid).Where("order_id = ?", orderId).Update("review_remark = ?", dis).Update("review_remark", now).Error
	if err != nil {
		return err
	}
	return nil
}

// ConvertToInvoiceList 将 model.InvoiceApplication 转换为 user_enter.InvoiceList
func (a *InvoiceApplication) ConvertToInvoiceList(application *InvoiceApplication) user_enter.InvoiceList {
	applicationTime := application.ApplicationTime.Format("20060102150405") // 申请时间
	reviewTime := application.ReviewTime.Format("20060102150405")           // 审核时间
	m := UserEnter{}
	merchant, _ := m.GetStatusById(application.MerId)

	return user_enter.InvoiceList{
		UserId:                       application.UserId,
		OrderId:                      application.OrderId,
		InvoiceType:                  application.InvoiceType,
		InvoiceTitle:                 application.InvoiceTitle,
		TaxpayerIdentificationNumber: application.TaxpayerIdentificationNumber,
		InvoiceAmount:                float32(application.InvoiceAmount),
		ApplicationTime:              applicationTime,
		ApplicationStatus:            int64(application.ApplicationStatus),
		ReviewTime:                   reviewTime,
		Type:                         application.Type,
		MerName:                      merchant.MerchantName,
	}
}
