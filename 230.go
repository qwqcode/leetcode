/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	node := root

	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		// 出栈
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 判断
		k--
		if k == 0 {
			return node.Val
		}

		node = node.Right
	}
}