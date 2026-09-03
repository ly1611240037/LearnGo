package math

import (
	"fmt"
	"testing"
)

func BenchmarkSumSerial(b *testing.B) {
	benchmarkSizes := []int{10, 1000, 100000}
	for _, size := range benchmarkSizes {
		data := makeData(size)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for b.Loop() {
				Sum(data)
			}
		})
	}
}

func BenchmarkSumParallel(b *testing.B) {
	benchmarkSizes := []int{10, 1000, 100000}
	for _, size := range benchmarkSizes {
		data := makeData(size)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for b.Loop() {
				SumParallel(data)
			}
		})
	}
}

func makeData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	return data
}
