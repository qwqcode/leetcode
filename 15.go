func threeSum(nums []int) [][]int {
	ans := [][]int{}

	sort.Ints(nums)

	// a + b + c == 0
	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break // 如果 a 大于 0，那么肯定凑不出三元组，因为没有比 a 小的负数来减了
		}
		// 去重 a
		if i > 0 && nums[i] == nums[i-1] { // -1 -1 -1 1 1 1... 的情况
			continue
		}

		l := i+1 // 注意是 i+1
		r := len(nums) - 1

		for r > l {
			a := nums[i]
			b := nums[l]
			c := nums[r]

			if a + b + c == 0 {
				// 收集结果
				ans = append(ans, []int{a, b, c})

				// 去重
				for l < r && nums[l] == b { // 跳过 x x x 1 1 1 2 2 2 x x x 的情况
					l++
				}
				for r > l && nums[r] == c {
					r--
				}
			} else if a + b + c > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return ans
}