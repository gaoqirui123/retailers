package request

// 生成邀请码
type GenerateInvitationCode struct {
	Type int64 `json:"type" binding:"required"  form:"type"` //邀请码类型
}

// 用户填写邀请码
type UserFillsInInvitationCode struct {
	Str string `json:"str" binding:"required"  form:"str"` //邀请码类型
}

// 分销等级设置
type DistributionLevelSetting struct {
	Img       string  `json:"img" binding:"required"  form:"img"`               //图片
	LevelName string  `json:"level_name" binding:"required"  form:"level_name"` //等级名称
	Level     int64   `json:"level" binding:"required"  form:"level"`           //等级
	One       float64 `json:"one" binding:"required"  form:"one"`               //一级返佣比例
	Two       float64 `json:"two" binding:"required"  form:"two"`               //二级返佣比例

}
