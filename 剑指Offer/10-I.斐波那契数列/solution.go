package leetcode

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		temp := b
		b = a%1000000007 + b%1000000007
		a = temp
	}
	return b % 1000000007
}
