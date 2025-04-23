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

		//中间件
		u.Use(pkg.JWTAuth("retailers"))
		u.GET("/detail", handler.UserDetail)                       //todo: 显示个人资料
		u.POST("/improve", handler.ImproveUser)                    //todo: 完善用户信息
		u.POST("/updatePassword", handler.UpdatePassWord)          //todo: 修改密码
		u.GET("/userLevel/list", handler.UserLevelList)            //todo: 会员页面展示
		u.GET("/userLevel/power/list", handler.UserLevelPowerList) //todo: 会员权益页面展示
		u.POST("/group/buy", handler.GroupBuying)                  //todo: 用户发起拼团
		u.POST("/add/usePower", handler.AddUsePower)               //todo: 用户使用权益

	}
}
