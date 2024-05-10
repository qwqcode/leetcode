func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	// dp 初始化：因为题目说到达 0 和 1 阶都不消耗体力
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ { // 遍历的顺序是从 2 开始到 n
		// 状态转移方程为：F(n) = min{ F(n - 1) + cost(n - 1), F(n - 2) + cost(n - 2) }
		dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
	}

	return dp[len(dp)-1] // dp[i] 的意思是：到达 i 位置花费的最小体力值为 dp[i]
}
