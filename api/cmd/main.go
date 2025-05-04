package main

import (
	"api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run("127.0.0.1:8080") // 监听并在 0.0.0.0:8080 上启动服务
	//pkg.FileMonitoring()
}
