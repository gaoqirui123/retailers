package model

import (
	"common/global"
)

type User struct {
	Uid            int64   `gorm:"column:uid;type:int UNSIGNED;comment:用户id;primaryKey;not null;" json:"uid"`                            // 用户id
	Account        string  `gorm:"column:account;type:varchar(32);comment:用户账号;not null;" json:"account"`                                // 用户账号
	Pwd            string  `gorm:"column:pwd;type:varchar(50);comment:用户密码;not null;" json:"pwd"`                                        // 用户密码
	RealName       string  `gorm:"column:real_name;type:varchar(25);comment:真实姓名;not null;" json:"real_name"`                            // 真实姓名
	Birthday       int64   `gorm:"column:birthday;type:int;comment:生日;not null;default:0;" json:"birthday"`                              // 生日
	CardId         string  `gorm:"column:card_id;type:varchar(20);comment:身份证号码;not null;" json:"card_id"`                               // 身份证号码
	Mark           string  `gorm:"column:mark;type:varchar(255);comment:用户备注;not null;" json:"mark"`                                     // 用户备注
	Nickname       string  `gorm:"column:nickname;type:varchar(60);comment:用户昵称;not null;" json:"nickname"`                              // 用户昵称
	Avatar         string  `gorm:"column:avatar;type:varchar(256);comment:用户头像;not null;" json:"avatar"`                                 // 用户头像
	Phone          string  `gorm:"column:phone;type:char(15);comment:手机号码;default:NULL;" json:"phone"`                                   // 手机号码
	Email          string  `gorm:"column:email;type:varchar(20);comment:邮箱;not null;" json:"email"`                                      // 邮箱
	FreezeAmount   float64 `gorm:"column:freeze_amount;type:decimal(8, 2);comment:冻结金额;default:NULL;" json:"freeze_amount"`              // 冻结金额
	NowMoney       float64 `gorm:"column:now_money;type:decimal(8, 2) UNSIGNED;comment:用户余额;not null;default:0.00;" json:"now_money"`    // 用户余额
	BrokeragePrice float64 `gorm:"column:brokerage_price;type:decimal(8, 2);comment:佣金金额;not null;default:0.00;" json:"brokerage_price"` // 佣金金额
	Integral       float64 `gorm:"column:integral;type:decimal(8, 2) UNSIGNED;comment:用户剩余积分;not null;default:0.00;" json:"integral"`    // 用户剩余积分
	SignNum        int64   `gorm:"column:sign_num;type:int;comment:连续签到天数;not null;default:0;" json:"sign_num"`                          // 连续签到天数
	Status         int64   `gorm:"column:status;type:tinyint(1);comment:1为正常，0为禁止;not null;default:1;" json:"status"`                    // 1为正常，0为禁止
	Level          int64   `gorm:"column:level;type:tinyint UNSIGNED;comment:等级;not null;default:0;" json:"level"`                       // 等级
	SpreadUid      int64   `gorm:"column:spread_uid;type:int UNSIGNED;comment:推广元id;not null;default:0;" json:"spread_uid"`              // 推广元id
	SpreadTime     int64   `gorm:"column:spread_time;type:int UNSIGNED;comment:推广员关联时间;not null;default:0;" json:"spread_time"`          // 推广员关联时间
	UserType       string  `gorm:"column:user_type;type:varchar(32);comment:用户类型;not null;" json:"user_type"`                            // 用户类型
	IsPromoter     int64   `gorm:"column:is_promoter;type:tinyint UNSIGNED;comment:是否为推广员/1是2否;not null;default:1;" json:"is_promoter"`  // 是否为推广员/1是2否
	PayCount       int64   `gorm:"column:pay_count;type:int UNSIGNED;comment:用户购买次数;default:0;" json:"pay_count"`                        // 用户购买次数
	SpreadCount    int64   `gorm:"column:spread_count;type:int;comment:下级人数;default:0;" json:"spread_count"`                             // 下级人数
	CleanTime      int64   `gorm:"column:clean_time;type:int;comment:清理会员时间;default:0;" json:"clean_time"`                               // 清理会员时间
	Address        string  `gorm:"column:address;type:varchar(255);comment:详细地址;not null;" json:"address"`                               // 详细地址
	Adminid        int64   `gorm:"column:adminid;type:int UNSIGNED;comment:管理员编号 ;default:0;" json:"adminid"`                            // 管理员编号
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
func (u *User) Detail(uid int64) (result []User, err error) {
	err = global.DB.Debug().Table("user").Where("uid=?", uid).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *User) FindId(id int) (result User, err error) {
	err = global.DB.Debug().Table("user").Where("uid=?", id).Find(&result).Error
	if err != nil {
		return User{}, err
	}
	return result, nil
}

func (u *User) Updated(id int64, users User) bool {
	err := global.DB.Debug().Table("user").Where("uid=?", id).Updates(users).Error
	if err != nil {
		return false
	}
	return true
}

func (u *User) UpdatedPassword(uid int64, password string) bool {
	err := global.DB.Debug().Table("user").Where("uid=?", uid).Limit(1).First(&u).Update("pwd", password).Error
	if err != nil {
		return false
	}
	return true
}

func (u *User) AddScore(score float64, uid int64) error {
	return global.DB.Debug().Table("user").Where("uid=?", uid).Update("integral", score).Error
}

// UpdatedSpreadUid 确认上级用户
func (u *User) UpdatedSpreadUid(uid int, sId string) bool {
	err := global.DB.Debug().Table("user").Where("uid=?", uid).Limit(1).First(&u).Update("spread_uid,level", sId).Error
	if err != nil {
		return false
	}
	return true
}

// 查找uid所在位置
func (u *User) FindDoneOrUpUid(id int64) (User, error) {
	var find User
	err := global.DB.Debug().Table("user").Where("uid = ?", id).Find(&find).Error
	if err != nil {
		return User{}, err
	}
	return find, nil
}

// 查找下级
func (u *User) FindDone(spreadUid int64) []User {
	var done []User
	err := global.DB.Debug().Table("user").Where("spread_uid = ?", spreadUid).Find(&done).Error
	if err != nil {
		return nil
	}
	return done
}

// 查找上级
func (u *User) FindUp(uid uint32) []User {
	var up []User
	err := global.DB.Debug().Table("user").Where("uid = ?", uid).Find(&up).Error
	if err != nil {
		return nil
	}
	return up
}

// 提现后，扣减账户余额
func (u *User) UpdateBalance(id int64, balance float64) error {
	err := global.DB.Debug().Table("user").Where("uid = ?", id).Update("now_money", balance).Error
	if err != nil {
		return nil
	}
	return nil
}
