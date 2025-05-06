package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/pkg"
	administrators "common/proto/admin"
	"github.com/gin-gonic/gin"
)

// AdminLogin TODO:管理员登录
func AdminLogin(c *gin.Context) {
	var data request.AdminLogin
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	login, err := client.AdminLogin(c, &administrators.AdminLoginReq{
		Account:  data.Account,
		Password: data.Pwd,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	claims := pkg.CustomClaims{
		ID: uint(login.AdminId),
	}
	token, err := pkg.NewJWT("2209A").CreateToken(claims)
	if err != nil {
		return
	}
	response.RespSuccess(c, "登录成功", token)
}

// ProcessEnter TODO:审核商家
func ProcessEnter(c *gin.Context) {
	var data request.ProcessEnter
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	ui := c.GetUint("userId")
	login, err := client.ProcessEnter(c, &administrators.ProcessEnterReq{
		AdminId:    int64(ui),
		MerchantId: data.MerchantId,
		Status:     data.Status,
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	if login.Greet == false {
		response.RespSuccess(c, "申请不合格，请重新申请", login)
	} else {
		response.RespSuccess(c, "审批成功", login)
	}
}
