package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	u := r.Group("/user")
	{
		u.POST("/login", handler.UserLogin)
		u.POST("/register", handler.UserRegister)
		//u.POST("/send", handler.SendSms)
		//u.POST("/forgot", handler.UserForgotPassWord)
		u.Use(pkg.JWTAuth("retailers"))
		u.GET("/detail", handler.UserDetail)              //显示个人资料
		u.POST("/improve", handler.ImproveUser)           //完善用户信息
		u.POST("/updatePassword", handler.UpdatePassWord) //修改密码
		//u.POST("/sign", handler.UserSign)
		//u.POST("/repair/sign", handler.UserRepairSign)
		//u.POST("/add/cart", handler.AddCart)
		//u.POST("/del/cart", handler.DeleteCart)
		//u.POST("/clear/cart", handler.ClearCart)
	}
}
