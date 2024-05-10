func uniquePaths(m int, n int) int {
	// m 行 n 列

    // 初始化 dp
    // 因为只能向下或向右走，所以第一行第一列的路径都分别只有 1 条
    dp := make([][]int, m)

    for i := 0; i < m; i++ {
    	dp[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
    	dp[i][0] = 1
    }
    for i := 0; i < n; i++ {
    	dp[0][i] = 1
    }

    // 遍历顺序：从 上往下、从左往右
    for i := 1; i < m; i++ {
    	for j := 1; j < n; j++ {
    		// 状态转移方程
    		dp[i][j] = dp[i-1][j] + dp[i][j-1]
    	}
    }

    return dp[m-1][n-1] // dp[i][j] 的意思是：走到 i 行 j 列共有 dp[i][j] 种走法
}