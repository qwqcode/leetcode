/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return nil
	}

	p := []*TreeNode{root}
	for len(p) > 0 {
		q := []*TreeNode{}
		v := 0

		for _, node := range p {
			if node == nil {
				continue
			}

			v = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		ans = append(ans, v)
		p = q
	}

	return ans
}