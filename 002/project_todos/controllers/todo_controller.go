package controllers

import (
	"net/http"
	"project_todos/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取所有待办事项
func GetTodos(c *gin.Context) {
	todos := models.GetAllTodos()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todos,
	})
}

// 获取单个待办事项
func GetTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "无效的ID",
		})
		return
	}

	todo, found := models.GetTodoById(uint(id))
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "待办事项不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todo,
	})
}

// 创建待办事项
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	newTodo := models.Create(todo)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   newTodo,
	})
}

// 更新待办事项
func UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "无效的ID",
		})
		return
	}

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	updatedTodo, found := models.Update(uint(id), todo)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "待办事项不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   updatedTodo,
	})
}

// 删除待办事项
func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "无效的ID",
		})
		return
	}

	found := models.DeleteTodo(uint(id))
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "待办事项不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   "待办事项已删除",
	})
}
