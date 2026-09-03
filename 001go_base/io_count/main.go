package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	text := string(data)
	charCount := utf8.RuneCountInString(text)
	wordCount := len(strings.Fields(text))
	lineCount := len(strings.Split(strings.TrimSuffix(text, "\n"), "\n"))

	fmt.Println("字符数:", charCount)
	fmt.Println("单词数:", wordCount)
	fmt.Println("行数:", lineCount)
}
