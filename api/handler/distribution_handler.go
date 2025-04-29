package handler

import (
	"api/client"
	"api/request"
	"api/response"
	"common/proto/distribution"
	"fmt"
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

// 分销等级设置
func DistributionLevelSetting(c *gin.Context) {

	var data request.DistributionLevelSetting
	if err := c.ShouldBind(&data); err != nil {
		response.RespError(c, fmt.Sprintf("参数错误%v", err))
		return
	}

	release, err := client.DistributionLevelSetting(c, &distribution.DistributionLevelSettingRequest{
		Img:       data.Img,
		LevelName: data.LevelName,
		Level:     data.Level,
		One:       float32(data.One),
		Two:       float32(data.Two),
	})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", release)
}

func TheCharts(c *gin.Context) {
	release, err := client.TheCharts(c, &distribution.TheChartsRequest{})
	if err != nil {
		response.RespError(c, err.Error())
		return
	}
	response.RespSuccess(c, "成功", release)
}
