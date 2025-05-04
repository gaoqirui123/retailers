package model

import "common/global"

type Order struct {
	Id                     int64   `gorm:"column:id;type:int UNSIGNED;comment:订单ID;primaryKey;not null;" json:"id"`                                                             // 订单ID
	OrderSn                string  `gorm:"column:order_sn;type:varchar(64);comment:订单号;not null;" json:"order_sn"`                                                              // 订单号
	Uid                    int64   `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                                                                      // 用户id
	RealName               string  `gorm:"column:real_name;type:varchar(32);comment:用户姓名;not null;" json:"real_name"`                                                           // 用户姓名
	UserPhone              string  `gorm:"column:user_phone;type:varchar(18);comment:用户电话;not null;" json:"user_phone"`                                                         // 用户电话
	UserAddress            string  `gorm:"column:user_address;type:varchar(100);comment:详细地址;not null;" json:"user_address"`                                                    // 详细地址
	CartId                 int64   `gorm:"column:cart_id;type:int;comment:购物车id;default:NULL;" json:"cart_id"`                                                                  // 购物车id
	FreightPrice           float64 `gorm:"column:freight_price;type:decimal(8, 2);comment:运费金额;default:0.00;" json:"freight_price"`                                             // 运费金额
	TotalNum               int64   `gorm:"column:total_num;type:int UNSIGNED;comment:订单商品总数;default:0;" json:"total_num"`                                                       // 订单商品总数
	TotalPrice             float64 `gorm:"column:total_price;type:decimal(8, 2) UNSIGNED;comment:订单总价;default:0.00;" json:"total_price"`                                        // 订单总价
	PayPrice               float64 `gorm:"column:pay_price;type:decimal(8, 2) UNSIGNED;comment:实际支付金额;default:0.00;" json:"pay_price"`                                          // 实际支付金额
	DeductionPrice         float64 `gorm:"column:deduction_price;type:decimal(8, 2) UNSIGNED;comment:抵扣金额;default:0.00;" json:"deduction_price"`                                // 抵扣金额
	CouponId               int64   `gorm:"column:coupon_id;type:int UNSIGNED;comment:优惠券id;default:0;" json:"coupon_id"`                                                        // 优惠券id
	CouponPrice            float64 `gorm:"column:coupon_price;type:decimal(8, 2) UNSIGNED;comment:优惠券金额;default:0.00;" json:"coupon_price"`                                     // 优惠券金额
	Paid                   int64   `gorm:"column:paid;type:tinyint UNSIGNED;comment:支付状态（2：未支付；3：已支付）;default:0;" json:"paid"`                                                  // 支付状态（2：未支付；3：已支付）
	PayTime                string  `gorm:"column:pay_time;type:varchar(20);comment:支付时间;default:NULL;" json:"pay_time"`                                                         // 支付时间
	PayType                int64   `gorm:"column:pay_type;type:int;comment:''支付方式(0-微信,1-支付宝,2-银行卡)'';default:NULL;" json:"pay_type"`                                           // ''支付方式(0-微信,1-支付宝,2-银行卡)''
	AddTime                int64   `gorm:"column:add_time;type:int UNSIGNED;comment:创建时间;default:NULL;" json:"add_time"`                                                        // 创建时间
	Status                 int64   `gorm:"column:status;type:tinyint(1);comment:订单状态（-1 : 申请仅退款； -2 : 申请退货退款成功 ；4：待发货；5：待收货；6：已收货；7：待评价；8：已退款；9：已取消）;default:0;" json:"status"` // 订单状态（-1 : 申请仅退款； -2 : 申请退货退款成功 ；4：待发货；5：待收货；6：已收货；7：待评价；8：已退款）
	RefundStatus           int64   `gorm:"column:refund_status;type:tinyint UNSIGNED;comment:0 未退款 1 申请中 2 已退款;default:0;" json:"refund_status"`                                // 0 未退款 1 申请中 2 已退款
	RefundReasonWapImg     string  `gorm:"column:refund_reason_wap_img;type:varchar(255);comment:退款图片;default:NULL;" json:"refund_reason_wap_img"`                              // 退款图片
	RefundReasonWapExplain string  `gorm:"column:refund_reason_wap_explain;type:varchar(255);comment:退款用户说明;default:NULL;" json:"refund_reason_wap_explain"`                    // 退款用户说明
	RefundReasonTime       int64   `gorm:"column:refund_reason_time;type:int UNSIGNED;comment:退款时间;default:NULL;" json:"refund_reason_time"`                                    // 退款时间
	RefundReasonWap        string  `gorm:"column:refund_reason_wap;type:varchar(255);comment:前台退款原因;default:NULL;" json:"refund_reason_wap"`                                    // 前台退款原因
	RefundReason           string  `gorm:"column:refund_reason;type:varchar(255);comment:不退款的理由;default:NULL;" json:"refund_reason"`                                            // 不退款的理由
	RefundPrice            float64 `gorm:"column:refund_price;type:decimal(8, 2) UNSIGNED;comment:退款金额;default:0.00;" json:"refund_price"`                                      // 退款金额
	DeliveryName           string  `gorm:"column:delivery_name;type:varchar(64);comment:快递名称/送货人姓名;default:NULL;" json:"delivery_name"`                                         // 快递名称/送货人姓名
	DeliveryType           string  `gorm:"column:delivery_type;type:varchar(32);comment:发货类型;default:NULL;" json:"delivery_type"`                                               // 发货类型
	DeliveryId             string  `gorm:"column:delivery_id;type:varchar(64);comment:快递单号/手机号;default:NULL;" json:"delivery_id"`                                               // 快递单号/手机号
	GainIntegral           int64   `gorm:"column:gain_integral;type:int UNSIGNED;comment:消费赚取积分;default:0;" json:"gain_integral"`                                               // 消费赚取积分
	UseIntegral            float64 `gorm:"column:use_integral;type:decimal(8, 2) UNSIGNED;comment:使用积分;default:0.00;" json:"use_integral"`                                      // 使用积分
	BackIntegral           float64 `gorm:"column:back_integral;type:decimal(8, 2) UNSIGNED;comment:给用户退了多少积分;default:NULL;" json:"back_integral"`                               // 给用户退了多少积分
	Mark                   string  `gorm:"column:mark;type:varchar(512);comment:备注;default:NULL;" json:"mark"`                                                                  // 备注
	IsDel                  int64   `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;default:0;" json:"is_del"`                                                           // 是否删除
	Unique                 string  `gorm:"column:unique;type:char(32);comment:唯一id(md5加密)类似id;default:NULL;" json:"unique"`                                                     // 唯一id(md5加密)类似id
	Remark                 string  `gorm:"column:remark;type:varchar(512);comment:管理员备注;default:NULL;" json:"remark"`                                                           // 管理员备注
	MerId                  int64   `gorm:"column:mer_id;type:int UNSIGNED;comment:商户ID;default:0;" json:"mer_id"`                                                               // 商户ID
	IsMerCheck             int64   `gorm:"column:is_mer_check;type:tinyint UNSIGNED;default:0;" json:"is_mer_check"`
	CombinationId          int64   `gorm:"column:combination_id;type:int UNSIGNED;comment:拼团商品id0一般商品;default:0;" json:"combination_id"`    // 拼团商品id0一般商品
	PinkId                 int64   `gorm:"column:pink_id;type:int UNSIGNED;comment:拼团id 0没有拼团;default:0;" json:"pink_id"`                   // 拼团id 0没有拼团
	Cost                   float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本价;default:NULL;" json:"cost"`                   // 成本价
	SeckillId              int64   `gorm:"column:seckill_id;type:int UNSIGNED;comment:秒杀商品ID;default:0;" json:"seckill_id"`                 // 秒杀商品ID
	BargainId              int64   `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价id;default:0;" json:"bargain_id"`                   // 砍价id
	VerifyCode             string  `gorm:"column:verify_code;type:varchar(12);comment:核销码;" json:"verify_code"`                             // 核销码
	StoreId                int64   `gorm:"column:store_id;type:int;comment:门店id;default:0;" json:"store_id"`                                // 门店id
	ShippingType           int64   `gorm:"column:shipping_type;type:tinyint(1);comment:配送方式 1=快递 ，2=门店自提;default:1;" json:"shipping_type"`  // 配送方式 1=快递 ，2=门店自提
	ClerkId                int64   `gorm:"column:clerk_id;type:int;comment:店员id;default:0;" json:"clerk_id"`                                // 店员id
	IsChannel              int64   `gorm:"column:is_channel;type:tinyint UNSIGNED;comment:支付渠道(0微信公众号1微信小程序);default:0;" json:"is_channel"` // 支付渠道(0微信公众号1微信小程序)
	IsRemind               int64   `gorm:"column:is_remind;type:tinyint UNSIGNED;comment:消息提醒;default:0;" json:"is_remind"`                 // 消息提醒
	IsSystemDel            int64   `gorm:"column:is_system_del;type:tinyint(1);comment:后台是否删除;default:0;" json:"is_system_del"`             // 后台是否删除
}

func (o *Order) AddOrder() error {
	return global.DB.Debug().Table("order").Create(&o).Error
}

func (o *Order) GetOrderStatus(sn string) error {
	return global.DB.Debug().Table("order").Where("order_sn = ?", sn).Limit(1).Find(&o).Error
}
func (o *Order) GetOrderSnUserId(sn string) Order {
	var order Order
	err := global.DB.Debug().Table("order").Where("order_sn = ?", sn).Limit(1).Find(&order).Error
	if err != nil {
		return order
	}
	return order

}

func (o *Order) UpdateOrderStatus(orderSn string, status int) error {
	return global.DB.Debug().Table("order").Where("order_sn = ?", orderSn).Limit(1).Update("paid", status).Error
}

func (o *Order) AddOrderPayTime(orderSn string, timeData string) error {
	return global.DB.Debug().Table("order").Where("order_sn = ?", orderSn).Limit(1).Update("pay_time", timeData).Error
}

func (o *Order) GetOrderPayList(userId, status int64) (list []*Order, err error) {
	err = global.DB.Debug().Table("order").Where("uid = ? and paid = ? and is_del = ?", userId, status, 0).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (o *Order) GetOrderStatusList(userId int64, status int64) (list []*Order, err error) {
	err = global.DB.Debug().Table("order").Where("uid = ? and status = ? and is_del = ?", userId, status, 0).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (o *Order) AllOrderList(userId int64) (list []*Order, err error) {
	err = global.DB.Debug().Table("order").Where("uid = ?", userId).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (o *Order) GetOrderDelList(userId int64, isDel int64) (list []*Order, err error) {
	err = global.DB.Debug().Table("order").Where("uid = ? and is_del = ?", userId, isDel).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (o *Order) GetOrderIdBy(userId int64, orderId int64) (list *Order, err error) {
	err = global.DB.Debug().Table("order").Where("uid = ? and id = ?", userId, orderId).Find(&o).Error
	if err != nil {
		return nil, err
	}
	return
}

// FindUserOrder 根据用户ID和订单ID查找订单
func (o *Order) FindUserOrder(uid, ueId int64) (*Order, error) {
	uo := Order{}
	order, err := uo.GetOrderIdBy(uid, ueId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) FindId(orderId int64) (result *Order, err error) {
	err = global.DB.Debug().Table("order").Where("id = ?", orderId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTotalOrderCount 获取订单总数
func (o *Order) GetTotalOrderCount() (int64, error) {
	var count int64
	err := global.DB.Debug().Table("order").Count(&count).Error
	return count, err
}

// GetTotalOrderAmount 获取订单总金额
func (o *Order) GetTotalOrderAmount() (float64, error) {
	var total float64
	err := global.DB.Debug().Table("order").Select("SUM(total_price)").Scan(&total).Error
	return total, err
}

// GetTotalRefundAmount 获取总退款数
func (o *Order) GetTotalRefundAmount() (float64, error) {
	var total float64
	err := global.DB.Debug().Table("order").Select("SUM(refund_price)").Scan(&total).Error
	return total, err
}
