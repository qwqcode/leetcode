// 动态规划的解法
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	dp := make([]int, n)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// 对动态规划解法进行状态压缩（滚动数组的思想 p、q、r 三个 int 变量）
func fib(n int) int {
	if n < 2 {
		return 0
	}
	sum := 1 // r
	p := 0
	q := 0

	for i := 2; i <= n; i++ {
		p = q
		q = sum
		sum = p + q
	}

	return sum
}