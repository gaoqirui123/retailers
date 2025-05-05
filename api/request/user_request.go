package request

type UserLogin struct {
	Account  string `json:"account" form:"account" binding:"required"`
	PassWord string `json:"passWord" form:"passWord" binding:"required"`
}

type UserRegister struct {
	Account  string `json:"account" form:"account" binding:"required"`
	PassWord string `json:"passWord" form:"passWord" binding:"required"`
	Pass     string `json:"pass" form:"pass" binding:"required"`
}

type ImproveUser struct {
	RealName string `form:"real_name" json:"real_name"  binding:"required"`
	Birthday int64  `form:"birthday" json:"birthday"  binding:"required"`
	CardId   string `form:"card_id" json:"card_id"  binding:"required"`
	Mark     string `form:"mark" json:"mark"`
	Nickname string `form:"nickname" json:"nickname"  binding:"required"`
	Avatar   string `form:"avatar" json:"avatar"  binding:"required"`
	Phone    string `form:"phone" json:"phone"  binding:"required"`
	Address  string `form:"address" json:"address"  binding:"required"`
}

type UpdatePassWord struct {
	NewPassword string `form:"new_password" json:"new_password"  binding:"required"`
}

type AddUserAddress struct {
	Province string `form:"province" xml:"province" json:"province"  binding:"required"`
	City     string `form:"city" xml:"city" json:"city"  binding:"required"`
	District string `form:"district" xml:"district" json:"district"  binding:"required"`
	Detail   string `form:"detail" xml:"detail" json:"detail"  binding:"required"`
}

type UserApplication struct {
	OrderId       int64   `form:"order_id" xml:"order_id" json:"order_id"  binding:"required"`
	InvoiceType   string  `form:"invoice_type" xml:"invoice_type" json:"invoice_type"  binding:"required"`
	InvoiceTitle  string  `form:"invoice_title" xml:"invoice_title" json:"invoice_title"  binding:"required"`
	InvoiceAmount float64 `form:"invoice_amount" xml:"invoice_amount" json:"invoice_amount"  binding:"required"`
	Type          string  `form:"type" xml:"type" json:"type"  binding:"required"`
}

<<<<<<< HEAD
type UpdatedAddress struct {
	Province string `form:"province" xml:"province" json:"province"  binding:"required"`
	City     string `form:"city" xml:"city" json:"city"  binding:"required"`
	District string `form:"district" xml:"district" json:"district"  binding:"required"`
	Detail   string `form:"detail" xml:"detail" json:"detail"  binding:"required"`
	RealName string `form:"real_name" xml:"real_name" json:"real_name"  binding:"required"`
	Phone    string `form:"phone" xml:"phone" json:"phone"  binding:"required"`
=======
type UserSignIn struct {
	SignData string `form:"signData" xml:"signData" json:"signData"  binding:"required"`
}

type UserMakeupSignIn struct {
	SignData string `form:"signData" xml:"signData" json:"signData"  binding:"required"`
}

type UserReceiveCoupon struct {
	CouponId int64 `form:"couponId" xml:"couponId" json:"couponId"  binding:"required"`
}

// 用户提现
type UserWithdraw struct {
	Amount         float64 `form:"amount" xml:"amount" json:"amount"  binding:"required"`        // 提现金额
	WithdrawMethod string  `form:"withdraw_method" xml:"withdraw_method" json:"withdraw_method"` // 提现方式，例如："支付宝", "微信", "银行卡"
	AccountInfo    string  `form:"account_info" xml:"account_info" json:"account_info"`          // 提现账户信息，根据提现方式不同而不同，如支付宝账号、银行卡号等

>>>>>>> 9d8aefe8fd97695cf5bb389ba8ce8b1bed1b904b
}
