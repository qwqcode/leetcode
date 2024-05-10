func longestPalindrome(s string) string {
	n := len(s)

	if n < 2 {
		return s
	}

	maxL := 1
	start := 0

	dp := make([][]bool, n)
	// 长度为 1 的全部为 true
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	// 从 L=2 开始迭代
	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			j := i + L - 1

			// 如果 j 长度超过 L 则打断
			if j >= n {
				break
			}

			// s[i] != s[j] 则 cache dp[i][j] = false
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i <= 2 {
					// i 最小为 0; j 最小为 0 + L - 1（L 最小为 2）= 1
					// 1~2 个字符，直接为 true
					dp[i][j] = true
				} else {
					// 大于 2 个字符，则引用前者 cache
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 满足条件 dp[i][j] == ture && 字符串长度大于最长的
			if dp[i][j] && (j-i+1) > maxL {
				maxL = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+maxL]
}
