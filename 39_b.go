// candidates = [2, 3, 6, 7], target = 7
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := [][]int{}
	path := []int{}

	var dfs func(begin, target int)
	dfs = func(begin, target int) {
		if target < 0 {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i := begin; i < n; i++ {
			path = append(path, candidates[i])

			dfs(i, target - candidates[i])

			path = path[:len(path) - 1]
		}
	}
	dfs(0, target)
	return ans
}