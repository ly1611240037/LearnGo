package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run . <目录路径>")
		return
	}

	root := os.Args[1]
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := entry.Info()
		if err != nil {
			return err
		}

		fmt.Printf("%-40s 目录=%-5t 大小=%-8d 修改时间=%s\n",
			path, entry.IsDir(), info.Size(), info.ModTime().Format("2006-01-02 15:04:05"))
		return nil
	})
	if err != nil {
		fmt.Println("遍历目录错误:", err)
	}
}
