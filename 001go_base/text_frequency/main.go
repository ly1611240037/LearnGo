package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run . <目录路径>")
		return
	}

	frequency := make(map[string]int)
	root := os.Args[1]

	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("读取目录错误:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".txt" {
			continue
		}

		path := filepath.Join(root, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("读取文件错误:", err)
			continue
		}

		for _, word := range strings.Fields(strings.ToLower(string(data))) {
			frequency[word]++
		}
	}

	fmt.Println("单词频率:")
	for word, count := range frequency {
		fmt.Printf("%s: %d 次\n", word, count)
	}
}
