package math

import (
	"runtime"
	"sync"
)

// Sum 串行计算切片中所有数字的和。
func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// SumParallel 使用多个 Goroutine 分段计算切片的和。
func SumParallel(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	workers := runtime.GOMAXPROCS(0)
	if workers > len(numbers) {
		workers = len(numbers)
	}

	partial := make([]int, workers)
	var wg sync.WaitGroup
	chunkSize := (len(numbers) + workers - 1) / workers

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		if start >= end {
			continue
		}

		wg.Add(1)
		go func(index, start, end int) {
			defer wg.Done()
			partial[index] = Sum(numbers[start:end])
		}(i, start, end)
	}

	wg.Wait()
	return Sum(partial)
}
