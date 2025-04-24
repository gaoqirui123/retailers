package router

import (
	"api/handler"
	"common/pkg"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	u := r.Group("/user")
	{
		u.POST("/login", handler.UserLogin)       //todo: 用户登录
		u.POST("/register", handler.UserRegister) //todo: 用户注册
		//u.POST("/send", handler.SendSms)
		//u.POST("/forgot", handler.UserForgotPassWord)
		u.Use(pkg.JWTAuth("retailers"))
		u.GET("/detail", handler.UserDetail)                     //todo: 显示个人资料
		u.POST("/improve", handler.ImproveUser)                  //todo: 完善用户信息
		u.POST("/updatePassword", handler.UpdatePassWord)        //todo: 修改密码
		u.GET("/userLevelList", handler.UserLevelList)           //todo: 会员页面展示
		u.GET("/userLevelPowerList", handler.UserLevelPowerList) //todo: 会员权益页面展示
		//u.POST("/sign", handler.UserSign)
		//u.POST("/repair/sign", handler.UserRepairSign)
	}
}
