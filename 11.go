func maxArea(height []int) int {
	l := 0
	r := len(height) - 1

	ans := 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		ans = max(ans, area)
		
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}