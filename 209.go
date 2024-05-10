func minSubArrayLen(target int, nums []int) int {
	f := false
	ans := len(nums)

	l := 0
	r := 0
	sum := 0

	for r < len(nums) {
		sum += nums[r]
        for sum >= target {
			f = true
            ans = min(ans, r - l + 1)
            sum -= nums[l]
            l++
        }
		r++
	}

	if !f {
		return 0
	}
	
	return ans
}

func min(a, b int) int { if a < b { return a }; return b }