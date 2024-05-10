func permute(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	used := map[int]bool{}

	path := []int{}
	var dfs func (depth int)
	dfs = func (depth int) {
		if depth == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i, num := range nums {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, num)
			dfs(depth + 1)
			path = path[:len(path) - 1]
			used[i] = false
		}
	}

	dfs(0)

	return ans
}
