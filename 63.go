func uniquePathsWithObstacles(obs [][]int) int {
	if len(obs) == 0 || len(obs[0]) == 0 {
		return 0
	}

	m := len(obs)
	n := len(obs[0])

	// 初始化 dp 数组
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 只能向右走或向下走
	for i := 0; i < m && obs[i][0] == 0; i++ { // 加了一个限定条件，遇到障碍则停止迭代
		dp[i][0] = 1
	}

	for i := 0; i < n && obs[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	// 遍历顺序：从上往下，从左往右，从 1 开始
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obs[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1] // dp[i][j] 表示走到 i 行 j 列总共有 dp[i][j] 种不同的路径
}