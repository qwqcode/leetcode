func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	ans := 0

	for len(queue) > 0 {
		for t := len(queue); t > 0; t-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans++
	}

	return ans
}