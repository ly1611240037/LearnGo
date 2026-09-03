package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("main 函数执行耗时：", time.Since(start).Nanoseconds(), "纳秒")
	}()

	time.Sleep(time.Second)
}
