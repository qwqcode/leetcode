func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	n := len(nums)
	l, r := 0, n - 1

	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
       		// 4 5 6 [7] 0 1 2 3
			// 如果第一个数字比 mid 小，则 左侧是升序，右侧是循环
			// 2. 判断 target 是在升序数组里，还是在循环数组里
			if nums[l] <= target && nums[mid] > target {
				// 如果目标比左边的大，并且比中间的小
				r = mid - 1 // 则在左侧升序数组里
			} else {
				// 否则在右侧循环数组里
				l = mid + 1
			}
		} else {
			// 否则左侧是循环，右侧是升序
			// 5 6 7 [0] 1 2 3 4
			if nums[r] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}