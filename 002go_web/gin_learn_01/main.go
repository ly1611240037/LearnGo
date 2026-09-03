package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建默认的路由引擎
	r := gin.Default()

	// 定义一个 GET 请求处理
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	// 启动服务器，监听 0.0.0.0:8080
	r.Run(":8080")
}
