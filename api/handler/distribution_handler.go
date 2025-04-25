package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/distribution"
	"github.com/gin-gonic/gin"
)

func GenerateInvitationCode(c *gin.Context) {

	var data request.GenerateInvitationCode
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	userId := c.GetUint("userId")
	release, err := client.GenerateInvitationCode(c, &distribution.GenerateInvitationCodeRequest{
		UserId: int64(userId),
		Type:   data.Type,
	})

	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", release)
}

func UserFillsInInvitationCode(c *gin.Context) {
	var data request.UserFillsInInvitationCode
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, "参数错误")
		return
	}
	userId := c.GetUint("userId")
	release, err := client.UserFillsInInvitationCode(c, &distribution.UserFillsInInvitationCodeRequest{
		UserId: uint32(userId),
		Str:    data.Str,
	})

	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", release)
}
