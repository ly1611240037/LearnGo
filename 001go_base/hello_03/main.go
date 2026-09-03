// main.go
package main

import "github.com/gin-gonic/gin" // 新依赖

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})
	r.Run()
}
