func productExceptSelf(nums []int) []int {
	n := len(nums)

	L := make([]int, n)
	L[0] = 1

	for i := 1; i < n; i++ {
		L[i] = nums[i - 1] * L[i - 1]
	}

	R := make([]int, n)
	R[n - 1] = 1
	for i := n - 2; i >= 0; i-- {
		R[i] = nums[i + 1] * R[i + 1]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = L[i] * R[i]
	}
	return ans
}