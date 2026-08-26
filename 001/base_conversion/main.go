package main

import (
	"fmt"
	"strconv"
)

func convertBases(decimal int) (string, string, string) {
	value := int64(decimal)
	return strconv.FormatInt(value, 2),
		strconv.FormatInt(value, 8),
		strconv.FormatInt(value, 16)
}

func main() {
	// 输入十进制数。
	fmt.Print("请输入一个十进制数: ")
	var decimal int
	if _, err := fmt.Scan(&decimal); err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	// 调用转换函数。
	binary, octal, hexadecimal := convertBases(decimal)

	// 输出转换结果。
	fmt.Println("二进制:", binary)
	fmt.Println("八进制:", octal)
	fmt.Println("十六进制:", hexadecimal)
}
