package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run . <文件路径>")
		return
	}

	sourcePath := os.Args[1]
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	ext := filepath.Ext(sourcePath)
	base := strings.TrimSuffix(filepath.Base(sourcePath), ext)
	backupName := fmt.Sprintf("%s_%s%s", base, time.Now().Format("20060102_150405"), ext)
	backupPath := filepath.Join(filepath.Dir(sourcePath), backupName)

	if err := os.WriteFile(backupPath, data, 0644); err != nil {
		fmt.Println("创建备份错误:", err)
		return
	}

	fmt.Println("备份完成:", backupPath)
}
