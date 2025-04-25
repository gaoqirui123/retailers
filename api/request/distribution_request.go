package request

// 生成邀请码
type GenerateInvitationCode struct {
	Type int64 `json:"type" binding:"required"  form:"type"` //邀请码类型
}

// 用户填写邀请码
type UserFillsInInvitationCode struct {
	Str string `json:"str" binding:"required"  form:"str"` //邀请码类型
}
