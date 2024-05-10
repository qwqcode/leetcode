/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	q := []*TreeNode{root}
	
	cnt := 0
	for len(q) > 0 {
		t := []*TreeNode{}

		for _, node := range q {
			if node == nil {
				continue
			}

			cnt++

			if node.Left != nil {
				t = append(t, node.Left)
			}
			if node.Right != nil {
				t = append(t, node.Right)
			}
		}

		q = t
	}

	return cnt
}
