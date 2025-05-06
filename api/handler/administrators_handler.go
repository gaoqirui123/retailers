package handler

import (
	"api/client"
	"api/request"
	"api/response"
	administrators "common/proto/admin"
	"github.com/gin-gonic/gin"
)

// AdminLogin TODO:管理员登录
func AdminLogin(c *gin.Context) {
	var data request.AdminLogin
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	login, err := client.AdminLogin(c, &administrators.AdminLoginReq{
		Account:  data.Account,
		Password: data.Pwd,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "登录成功", login)
}

// ProcessEnter TODO:审核商家
func ProcessEnter(c *gin.Context) {
	var data request.ProcessEnter
	err := c.ShouldBind(&data)
	if err != nil {
		response.RespError(c, 201, err.Error())
		return
	}
	ui := c.GetUint("userId")
	login, err := client.ProcessEnter(c, &administrators.ProcessEnterReq{
		AdminId:    int64(ui),
		MerchantId: data.MerchantId,
		Status:     data.Status,
	})
	if err != nil {
		response.RespError(c, 500, err.Error())
		return
	}
	response.RespSuccess(c, 200, "审批成功", login)
}
