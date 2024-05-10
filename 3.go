func lengthOfLongestSubstring(s string) int {
    win := make(map[byte]bool)

	l := 0
	ans := 0

    for r := 0; r < len(s); r++ {
		if win[s[r]] {
			l++
			win[s[l]] = false
		}

		win[s[r]] = true
		ans = max(ans, r - l + 1)
    }

	return ans
}

func max(a, b int) int { if a > b { return a }; return b }