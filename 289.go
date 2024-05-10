func gameOfLife(board [][]int)  {
	neighbors := []int{0, 1, -1}
	
	rowNum := len(board)
	colNum := len(board[0])

	// 创建复制数组 boardCopy
	boardCopy := make([][]int, rowNum)

	for i := range boardCopy {
		boardCopy[i] = make([]int, colNum)
	}

	// 从原数组复制一份到 boardCopy 中
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			boardCopy[row][col] = board[row][col]
		}
	}

	// 遍历面板每一个格子的细胞
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			// 对于每一个细胞统计其八个相邻位置里的活细胞数量
			liveNeighbors := 0

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if neighbors[i] == 0 && neighbors[j] == 0 {
						continue
					}

					r := row + neighbors[i]
					c := col + neighbors[j]

					// 查看相邻的细胞是否是活细胞
					if (r < rowNum && r >= 0) && (c < colNum && c >= 0) && boardCopy[r][c] == 1 {
						liveNeighbors += 1
					}
				}
			}

			// 规则 1 或规则 3
			if (boardCopy[row][col] == 1) && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[row][col] = 0
			}
			if boardCopy[row][col] == 0 && liveNeighbors == 3 {
				board[row][col] = 1
			}
		}
	}
}