func zigzagLevelOrder(root *TreeNode) [][]int {
	p := []*TreeNode{root}
	ans := [][]int{}

	i := 0
	for len(p) > 0 {
		q := []*TreeNode{}
		data := []int{}

		for _, node := range p {
			if node == nil {
				continue
			}

			if i % 2 != 0 {
				data = append([]int{node.Val}, data...)
			} else {
				data = append(data, node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			} 
		}

		if len(data) > 0 {
			ans = append(ans, data)
		}
		p = q
		i++
	}

	return ans
}