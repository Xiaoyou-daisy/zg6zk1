package main

import (
	"github.com/gin-gonic/gin"
	"zg6zk1/apiway/routers"
)

func main() {
	// 创建一个默认的 gin 路由
	router := gin.Default()

	// 注册登录路由
	routers.LoginRouter()

	// 定义一个 GET 请求的路由，当访问 /ping 时，返回一个 JSON 响应
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
