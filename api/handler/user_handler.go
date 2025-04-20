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
