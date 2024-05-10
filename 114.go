/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	list := preOrderTraversal(root)

	n := len(list)

	for i := 1; i < n; i++ {
		prev, curt := list[i - 1], list[i]
		prev.Left, prev.Right = nil, curt
	}
}

func preOrderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preOrderTraversal(root.Left)...)
		list = append(list, preOrderTraversal(root.Right)...)
	}
	return list
}