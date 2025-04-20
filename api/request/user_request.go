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
