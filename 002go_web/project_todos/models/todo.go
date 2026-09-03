package models

import "time"

type Todo struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreateAt    time.Time `json:"create_at"`
}

var todos = []Todo{
	{ID: 1, Title: "学习gin框架", Description: "学习gin框架的路由，中间件等知识", Completed: false, CreateAt: time.Now()},
	{ID: 2, Title: "复习C++相关知识", Description: "复习C++基本语法。数据结构等", Completed: false, CreateAt: time.Now()},
}

func GetAllTodos() []Todo {
	return todos
}

func GetTodoById(id uint) (Todo, bool) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, true
		}
	}
	return Todo{}, false
}

func Create(todo Todo) Todo {
	todo.ID = uint(len(todos) + 1)
	todo.CreateAt = time.Now()
	todos = append(todos, todo)
	return todo
}

func Update(id uint, updateTodo Todo) (Todo, bool) {
	for i, todo := range todos {
		if todo.ID == id {
			updateTodo.ID = id
			updateTodo.CreateAt = todo.CreateAt
			todos[i] = updateTodo
			return updateTodo, true
		}
	}
	return Todo{}, false
}

func DeleteTodo(id uint) bool {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
