package main

import (
	"project_todos/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", controllers.GetTodos)
			todos.GET("/:id", controllers.GetTodo)
			todos.POST("", controllers.CreateTodo)
			todos.PUT("/:id", controllers.UpdateTodo)
			todos.DELETE("/:id", controllers.DeleteTodo)
		}
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
