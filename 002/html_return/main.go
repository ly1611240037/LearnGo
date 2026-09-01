package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title":   "标题",
			"message": "你好世界",
		})
	})

	r.Run(":8080")
}
