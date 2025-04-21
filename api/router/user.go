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
		u.GET("/detail", handler.UserDetail)
		//u.POST("/update", handler.UserUpdatePassWord)
		//u.POST("/improve", handler.UserImprove)
		//u.POST("/sign", handler.UserSign)
		//u.POST("/repair/sign", handler.UserRepairSign)
		//u.POST("/add/cart", handler.AddCart)
		//u.POST("/del/cart", handler.DeleteCart)
		//u.POST("/clear/cart", handler.ClearCart)
	}
}
