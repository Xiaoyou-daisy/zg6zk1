package routers

import (
	"github.com/gin-gonic/gin"
	"zg6zk1/apiway/basic/pkg"
	"zg6zk1/apiway/handler/apihandler"
)

func LoginRouter() {
	router := gin.Default()

	// 简单的路由组: v1
	{
		v1 := router.Group("/v1")
		v1.POST("/login", apihandler.Login)
	}
	{
		v2 := router.Group("/v2")
		v2.Use(pkg.JWTAuth("zk"))
		v2.POST("/list", apihandler.List)
	}
	router.Run(":8080")
}
