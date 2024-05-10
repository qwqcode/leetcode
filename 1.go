// 暴力枚举求解法
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 哈希表法
func twoSumHash(nums []int, target int) []int {
	m := map[int]int{}

	for i, x := range nums {
		if p, ok := m[target-x]; ok {
			return []int{i, p}
		}
		m[x] = i
	}
	return nil
}

// 双指针法
func twoSum2(nums []int, target int) []int {
	sort.Ints(nums) // 注意如果是返回下标，排序过后下标会发生变化

	n := len(nums)

	l := 0
	r := n - 1

	for l < r {
		x := nums[l] + nums[r]
		if x == target {
			return []int{l, r}
		} else if x > target {
			r--
		} else if x < target {
			l--
		}
	}
	return []int{}
}
