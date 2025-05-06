package response

import (
	"github.com/gin-gonic/gin"
)

type CodeStatus int

// 定义常见的状态码常量
const (
	StatusOK              CodeStatus = 200 // Todo: 请求成功
	StatusSqlErr          CodeStatus = 201 // Todo: 参数错误
	StatusBadRequest      CodeStatus = 400 // Todo: 无效的请求
	StatusUnauthorized    CodeStatus = 401 // Todo: 未授权
	StatusPaymentRequired CodeStatus = 402 // Todo: 授权已过期
	StatusForbidden       CodeStatus = 403 // Todo: 权限不足
	StatusNotFound        CodeStatus = 404 // Todo: 资源未找到，路由配置错误
	StatusInternalError   CodeStatus = 500 // Todo: 服务端报错提示，数据验证错误
)

// 定义状态码对应的消息映射
var statusMessages = map[CodeStatus]string{
	StatusOK:              "请求成功",
	StatusSqlErr:          "参数错误",
	StatusBadRequest:      "无效的请求",
	StatusUnauthorized:    "未授权",
	StatusPaymentRequired: "授权已过期",
	StatusForbidden:       "权限不足",
	StatusNotFound:        "资源未找到，路由配置错误",
	StatusInternalError:   "服务端报错提示，数据验证错误",
}

// GetMessage 根据状态码获取对应的消息
func GetMessage(code CodeStatus) string {
	if msg, ok := statusMessages[code]; ok {
		return msg
	}
	return "未知状态码"
}

// RespError Api统一错误响应函数，支持多种错误状态码
func RespError(c *gin.Context, code CodeStatus, msg string) {
	resp := map[string]interface{}{
		"code": code,
		"msg":  GetMessage(code) + msg,
		"data": nil,
	}
	c.JSON(int(code), resp)
}

// RespSuccess Api统一成功响应函数
func RespSuccess(c *gin.Context, code CodeStatus, msg string, data interface{}) {
	resp := map[string]interface{}{
		"code": code,
		"msg":  GetMessage(code) + msg,
		"data": data,
	}
	c.JSON(int(code), resp)
}
