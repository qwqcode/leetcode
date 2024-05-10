func trap(height []int) int {
	ans := 0

	n := len(height)
	if n == 0 {
		return 0
	}
	leftHeight := make([]int, n)
	rightHeight := make([]int, n)
	
	leftHeight[0] = height[0]
	rightHeight[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftHeight[i] = max(leftHeight[i - 1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightHeight[i] = max(rightHeight[i + 1], height[i])
	}

	for i, h := range height {
		ans += min(leftHeight[i], rightHeight[i]) - h
	}

	return ans
}