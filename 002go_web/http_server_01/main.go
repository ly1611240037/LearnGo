package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const todoURL = "https://jsonplaceholder.typicode.com/todos"

func main() {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println("加载 Todo 数据失败:", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/todo/list", listHandler(todos))
	mux.HandleFunc("/api/todo/detail/", detailHandler(todos))

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Println("HTTP 服务已启动: http://localhost:8080")
	fmt.Println("列表接口: http://localhost:8080/api/todo/list")
	fmt.Println("详情接口: http://localhost:8080/api/todo/detail/1")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("服务启动失败:", err)
	}
}

func loadTodos() ([]Todo, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(todoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("远程服务返回状态码 %d", resp.StatusCode)
	}

	var todos []Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func listHandler(todos []Todo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, "只支持 GET 请求", http.StatusMethodNotAllowed)
			return
		}
		writeJSON(w, http.StatusOK, todos)
	}
}

func detailHandler(todos []Todo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, "只支持 GET 请求", http.StatusMethodNotAllowed)
			return
		}

		idText := strings.TrimPrefix(r.URL.Path, "/api/todo/detail/")
		id, err := strconv.Atoi(idText)
		if err != nil {
			writeError(w, "todoid 必须是整数", http.StatusBadRequest)
			return
		}

		for _, todo := range todos {
			if todo.ID == id {
				writeJSON(w, http.StatusOK, todo)
				return
			}
		}

		writeError(w, "todo 不存在", http.StatusNotFound)
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, message string, status int) {
	writeJSON(w, status, map[string]string{"error": message})
}
