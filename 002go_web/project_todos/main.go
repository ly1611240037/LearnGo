package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project_todos/controllers"
	"syscall"
	"time"

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
	router := setupRouter()

	svr := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		log.Fatal("服务器强制关闭: ", err)
	}

	log.Println("服务器优雅关闭")
}
