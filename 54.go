func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	ans := []int{}

	// 定义四种状态 (x, y)
	states := [][]int{ []int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1} }

	// 统计 matrix 长度
	rowLen := len(matrix)
	colLen := len(matrix[0])

	// 初始化记录已访问元素的表
	visited := make([][]bool, rowLen)
	for i := range visited {
		visited[i] = make([]bool, colLen)
	}

	// 遍历矩阵
	stateID := 0
	x := 0
	y := 0

	length := rowLen * colLen
	for i := 0; i < length; i++ {
		ans = append(ans, matrix[y][x])

		visited[y][x] = true

		// 如果触碰到边界则切换下一个状态
		nextX := x + states[stateID][0] // a.预先验证
		nextY := y + states[stateID][1]
		if nextX >= colLen || nextY >= rowLen || nextX < 0 || nextY < 0 || visited[nextY][nextX] {
			stateID = (stateID + 1) % 4 // 状态切换
		}

		x = x + states[stateID][0] // b.实际变更
		y = y + states[stateID][1]
	}

	return ans
}