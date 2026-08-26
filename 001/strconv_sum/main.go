package main

import (
	"fmt"
	"strconv"
)

func sumNumbers(numbers []string) (int, error) {
	sum := 0
	for _, text := range numbers {
		number, err := strconv.Atoi(text)
		if err != nil {
			return 0, fmt.Errorf("%q 不是有效的整数: %w", text, err)
		}
		sum += number
	}
	return sum, nil
}

func main() {
	// main 函数负责接收输入。
	fmt.Print("请输入数字个数: ")
	var count int
	if _, err := fmt.Scan(&count); err != nil || count < 1 {
		fmt.Println("请输入一个正整数")
		return
	}

	numbers := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Printf("请输入第 %d 个数字字符串: ", i+1)
		if _, err := fmt.Scan(&numbers[i]); err != nil {
			fmt.Println("读取输入失败:", err)
			return
		}
	}

	// 调用处理函数，然后输出结果。
	sum, err := sumNumbers(numbers)
	if err != nil {
		fmt.Println("转换错误:", err)
		return
	}
	fmt.Println("总和:", sum)
}
