package math

// 计算两个数的最大公约数
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
