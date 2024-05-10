func removeDuplicates(nums []int) int {
	// 快慢指针法

	n := len(nums)
	if n == 0 {
		return 0
	}

    fast := 1
	slow := 1

	for ; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}