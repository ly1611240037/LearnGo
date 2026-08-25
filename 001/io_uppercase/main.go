package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	upperText := strings.ToUpper(string(data))
	if err := os.WriteFile("output.txt", []byte(upperText), 0644); err != nil {
		fmt.Println("写入文件错误:", err)
		return
	}

	fmt.Println("转换完成，结果已写入 output.txt")
}
