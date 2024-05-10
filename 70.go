func climbStairs(n int) int {
    if n <= 0 {
        return 0
    }
    if n <= 2 {
        return n
    }

    // 爬楼梯问题，通过分析能够找到规律
    // 爬 1 层楼梯，有 1 种解法
    // 爬 2 层楼梯，有 2 种解法
    // 爬 3 层楼梯，有 3 种解法
    // 爬 4 层楼梯，有 5 种解法

    // 得到状态转移方程 f(n) = f(n-1) + f(n-2)

    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    dp[2] = 2

    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}