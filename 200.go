func dfs(grid [][]byte, int r, int c) {
	nr := len(grid)
	nc := len(grid[0])

	if r < 0 || c < 0 || r >= nr || c >= nr || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r - 1, c)
	dfs(grid, r + 1, c)
	dfs(grid, r, c - 1)
	dfs(grid, r, c + 1)
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	cnt := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				cnt++
				dfs(grid, r, c)
			}
		}
	}

	return cnt
}