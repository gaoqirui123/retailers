package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/pkg"
	"common/proto/user"
	"github.com/gin-gonic/gin"
)

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

func UserDetail(c *gin.Context) {
	userId := c.GetUint("userId")
	detail, err := client.UserDetail(c, &user.UserDetailRequest{Uid: int32(userId)})
	if err != nil {
		response.RespError(c, "查看失败")
		return
	}
	response.RespSuccess(c, "个人资料显示成功", detail)
}

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

func UserLevelList(c *gin.Context) {
	list, err := client.UserLevelList(c, &user.UserLevelListRequest{})
	if err != nil {
		response.RespError(c, "会员页面展示失败")
		return
	}
	response.RespSuccess(c, "会员页面展示成功", list)
}
