package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/pkg"
	"common/proto/user"
	"github.com/gin-gonic/gin"
)

// TODO: 用户登录
func UserLogin(c *gin.Context) {
	var data request.UserLogin
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	login, err := client.UserLogin(c, &user.UserLoginRequest{
		Account:  data.Account,
		PassWord: data.PassWord,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if login.UserId == 0 {
		response.RespError(c, "登录失败")
		return
	}
	token, _ := pkg.NewJWT("retailers").CreateToken(pkg.CustomClaims{
		ID: uint(login.UserId),
	})
	response.RespSuccess(c, "登录成功", token)
}

// TODO:用户注册
func UserRegister(c *gin.Context) {
	var data request.UserRegister
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	register, err := client.UserRegister(c, &user.UserRegisterRequest{
		Account:  data.Account,
		PassWord: data.PassWord,
		Pass:     data.Pass,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if register.UserId == 0 {
		response.RespError(c, "注册失败")
		return
	}
	response.RespSuccess(c, "注册成功", register.UserId)
}

// TODO:展示个人资料
func UserDetail(c *gin.Context) {
	userId := c.GetUint("userId")
	detail, err := client.UserDetail(c, &user.UserDetailRequest{Uid: int32(userId)})
	if err != nil {
		response.RespError(c, "查看失败")
		return
	}
	response.RespSuccess(c, "个人资料显示成功", detail)
}

// TODO:完善用户信息
func ImproveUser(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.ImproveUser
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}

	improveUser, err := client.ImproveUser(c, &user.ImproveUserRequest{
		RealName: data.RealName,
		Birthday: int32(data.Birthday),
		CardId:   data.CardId,
		Mark:     data.Mark,
		Nickname: data.Nickname,
		Avatar:   data.Avatar,
		Phone:    data.Phone,
		Address:  data.Address,
		Id:       int32(userId),
	})
	if err != nil {
		response.RespError(c, "用户完善信息失败")
		return
	}
	response.RespSuccess(c, "用户完善信息成功", improveUser)
}

// TODO:修改密码
func UpdatePassWord(c *gin.Context) {
	userId := c.GetUint("userId")
	var data request.UpdatePassWord
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}

	password, err := client.UpdatedPassword(c, &user.UpdatedPasswordRequest{
		Uid:         int32(userId),
		NewPassword: data.NewPassword,
	})
	if err != nil {
		response.RespError(c, "用户修改密码失败")
		return
	}
	response.RespSuccess(c, "用户修改密码成功", password)
}

// TODO:会员页面展示
func UserLevelList(c *gin.Context) {
	list, err := client.UserLevelList(c, &user.UserLevelListRequest{})
	if err != nil {
		response.RespError(c, "会员页面展示失败")
		return
	}
	response.RespSuccess(c, "会员页面展示成功", list)
}

// TODO:会员权益页面展示
func UserLevelPowerList(c *gin.Context) {
	list, err := client.UserLevelPowerList(c, &user.UserLevelPowerListRequest{})
	if err != nil {
		response.RespError(c, "会员权益页面展示失败")
		return
	}
	response.RespSuccess(c, "会员页面展示成功", list)
}

// TODO:用户使用权益
func AddUsePower(c *gin.Context) {
	userId := c.GetUint("userId")

	power, err := client.AddUsePower(c, &user.AddUsePowerRequest{
		Uid: int32(userId),
	})
	if err != nil {
		response.RespError(c, "用户使用权益失败")
		return
	}
	response.RespSuccess(c, "用户使用权益成功", power)
}

// TODO: 用户使用权益表展示
func UsePowerList(c *gin.Context) {
	list, err := client.UsePowerList(c, &user.UsePowerListRequest{})
	if err != nil {
		response.RespError(c, "用户使用权益表展示失败")
		return
	}
	response.RespSuccess(c, "用户使用权益表展示成功", list)
}
