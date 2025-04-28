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
		u.GET("/userLevel/list", handler.UserLevelList)            //todo: 会员页面展示
		u.GET("/userLevel/power/list", handler.UserLevelPowerList) //todo: 会员权益页面展示
		u.GET("/usePower/list", handler.UsePowerList)              //todo: 用户使用权益表展示

		//中间件
		u.Use(pkg.JWTAuth("retailers"))
		u.GET("/detail", handler.UserDetail)               //todo: 显示个人资料
		u.POST("/improve", handler.ImproveUser)            //todo: 完善用户信息
		u.POST("/updatePassword", handler.UpdatePassWord)  //todo: 修改密码
		u.POST("/add/usePower", handler.AddUsePower)       //todo: 用户使用权益
		u.POST("/add/text", handler.AddText)               //todo: 会员分添加记录
		u.POST("/add/userAddress", handler.AddUserAddress) //todo: 用户添加地址
		u.Use(pkg.JWTAuth("retailers"))
		u.POST("/sign", handler.UserSignIn)                 //todo: 用户签到
		u.POST("/makeup/sign", handler.UserMakeupSignIn)    //todo: 用户补签
		u.POST("/userApplication", handler.UserApplication) //todo: 用户申请发票
	}
}
