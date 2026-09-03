package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Todo 表示接口返回的一条待办事项。
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/todos", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("请求失败，状态码:", resp.StatusCode)
		return
	}

	var todos []Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		fmt.Println("解析响应失败:", err)
		return
	}

	for _, todo := range todos {
		fmt.Printf("title: %s, userId: %d, id: %d, completed: %t\n",
			todo.Title, todo.UserID, todo.ID, todo.Completed)
	}
}
