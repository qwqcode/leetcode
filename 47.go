func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	used := map[int]bool{}
	path := []int{}

	var dfs func(depth int)
	dfs = func(depth int) {
		if depth == len(nums) {
			// 记录排列结果
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i, num := range nums {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}

			path = append(path, num)
			used[i] = true
			dfs(depth + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	// 先对数组进行排序
	sort.Ints(nums)
	dfs(0)

	return ans
}