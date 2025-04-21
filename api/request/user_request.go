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
	Birthday int    `form:"birthday" json:"birthday"  binding:"required"`
	CardId   string `form:"card_id" json:"card_id"  binding:"required"`
	Mark     string `form:"mark" json:"mark"`
	Nickname string `form:"nickname" json:"nickname"  binding:"required"`
	Avatar   string `form:"avatar" json:"avatar"  binding:"required"`
	Phone    string `form:"phone" json:"phone"  binding:"required"`
	Address  string `form:"address" json:"address"  binding:"required"`
}
