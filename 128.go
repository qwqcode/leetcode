func longestConsecutive(nums []int) int {
	// 1. 去重
	numSet := map[int]bool{}
	for _, v := range nums {
		numSet[v] = true
	}

	// 2. 求最长
	ans := 0
	for n := range numSet {
		if numSet[n-1] {
			continue
		}

		t := 1
		curt := n
		for numSet[curt+1] {
			t++
			curt++
		}
		ans = max(ans, t)
	}

	return ans
}