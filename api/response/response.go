package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  msg,
		"data": nil,
	})
}

func RespSuccess(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}
