package model

import "common/global"

type User struct {
	Uid            uint32  `gorm:"column:uid;type:int UNSIGNED;comment:用户id;primaryKey;not null;" json:"uid"`                            // 用户id
	Account        string  `gorm:"column:account;type:varchar(32);comment:用户账号;not null;" json:"account"`                                // 用户账号
	Pwd            string  `gorm:"column:pwd;type:varchar(50);comment:用户密码;not null;" json:"pwd"`                                        // 用户密码
	RealName       string  `gorm:"column:real_name;type:varchar(25);comment:真实姓名;not null;" json:"real_name"`                            // 真实姓名
	Birthday       int32   `gorm:"column:birthday;type:int;comment:生日;not null;default:0;" json:"birthday"`                              // 生日
	CardId         string  `gorm:"column:card_id;type:varchar(20);comment:身份证号码;not null;" json:"card_id"`                               // 身份证号码
	Mark           string  `gorm:"column:mark;type:varchar(255);comment:用户备注;not null;" json:"mark"`                                     // 用户备注
	PartnerId      int32   `gorm:"column:partner_id;type:int;comment:合伙人id;not null;default:0;" json:"partner_id"`                       // 合伙人id
	GroupId        int32   `gorm:"column:group_id;type:int;comment:用户分组id;not null;default:0;" json:"group_id"`                          // 用户分组id
	Nickname       string  `gorm:"column:nickname;type:varchar(60);comment:用户昵称;not null;" json:"nickname"`                              // 用户昵称
	Avatar         string  `gorm:"column:avatar;type:varchar(256);comment:用户头像;not null;" json:"avatar"`                                 // 用户头像
	Phone          string  `gorm:"column:phone;type:char(15);comment:手机号码;default:NULL;" json:"phone"`                                   // 手机号码
	AddTime        uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;not null;default:0;" json:"add_time"`                   // 添加时间
	AddIp          string  `gorm:"column:add_ip;type:varchar(16);comment:添加ip;not null;" json:"add_ip"`                                  // 添加ip
	LastTime       uint32  `gorm:"column:last_time;type:int UNSIGNED;comment:最后一次登录时间;not null;default:0;" json:"last_time"`             // 最后一次登录时间
	LastIp         string  `gorm:"column:last_ip;type:varchar(16);comment:最后一次登录ip;not null;" json:"last_ip"`                            // 最后一次登录ip
	NowMoney       float64 `gorm:"column:now_money;type:decimal(8, 2) UNSIGNED;comment:用户余额;not null;default:0.00;" json:"now_money"`    // 用户余额
	BrokeragePrice float64 `gorm:"column:brokerage_price;type:decimal(8, 2);comment:佣金金额;not null;default:0.00;" json:"brokerage_price"` // 佣金金额
	Integral       float64 `gorm:"column:integral;type:decimal(8, 2) UNSIGNED;comment:用户剩余积分;not null;default:0.00;" json:"integral"`    // 用户剩余积分
	SignNum        int32   `gorm:"column:sign_num;type:int;comment:连续签到天数;not null;default:0;" json:"sign_num"`                          // 连续签到天数
	Status         int8    `gorm:"column:status;type:tinyint(1);comment:1为正常，0为禁止;not null;default:1;" json:"status"`                    // 1为正常，0为禁止
	Level          uint8   `gorm:"column:level;type:tinyint UNSIGNED;comment:等级;not null;default:0;" json:"level"`                       // 等级
	SpreadUid      uint32  `gorm:"column:spread_uid;type:int UNSIGNED;comment:推广元id;not null;default:0;" json:"spread_uid"`              // 推广元id
	SpreadTime     uint32  `gorm:"column:spread_time;type:int UNSIGNED;comment:推广员关联时间;not null;default:0;" json:"spread_time"`          // 推广员关联时间
	UserType       string  `gorm:"column:user_type;type:varchar(32);comment:用户类型;not null;" json:"user_type"`                            // 用户类型
	IsPromoter     uint8   `gorm:"column:is_promoter;type:tinyint UNSIGNED;comment:是否为推广员;not null;default:0;" json:"is_promoter"`       // 是否为推广员
	PayCount       uint32  `gorm:"column:pay_count;type:int UNSIGNED;comment:用户购买次数;default:0;" json:"pay_count"`                        // 用户购买次数
	SpreadCount    int32   `gorm:"column:spread_count;type:int;comment:下级人数;default:0;" json:"spread_count"`                             // 下级人数
	CleanTime      int32   `gorm:"column:clean_time;type:int;comment:清理会员时间;default:0;" json:"clean_time"`                               // 清理会员时间
	Addres         string  `gorm:"column:addres;type:varchar(255);comment:详细地址;not null;" json:"addres"`                                 // 详细地址
	Adminid        uint32  `gorm:"column:adminid;type:int UNSIGNED;comment:管理员编号 ;default:0;" json:"adminid"`                            // 管理员编号
	LoginType      string  `gorm:"column:login_type;type:varchar(36);comment:用户登陆类型，h5,wechat,routine;not null;" json:"login_type"`      // 用户登陆类型，h5,wechat,routine
}

func (u *User) UserLogin(account string) error {
	return global.DB.Debug().Table("user").Where("account = ?", account).Limit(1).Find(&u).Error
}

func (u *User) UserRegister() error {
	return global.DB.Debug().Table("user").Create(&u).Error
}

func (u *User) GetUserIdBy(uid int64) error {
	return global.DB.Debug().Table("user").Where("uid = ?", uid).Limit(1).Find(&u).Error
}
