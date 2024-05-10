// candidates = [2, 3, 6, 7], target = 7
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	comb := []int{}
	var dfs func(target, i int)
	dfs = func(target, i int) {
		if i == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		dfs(target, i+1)
		if target - candidates[i] >= 0 {
			comb = append(comb, candidates[i])
			dfs(target - candidates[i], i)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}