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
