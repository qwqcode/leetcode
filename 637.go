/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}

	if root == nil {
		return ans
	}
	
	p := []*TreeNode{root}

	for len(p) > 0 {
		q := []*TreeNode{}
		total := 0
		count := 0

		for _, node := range p {
			if node == nil {
				continue
			}

			total += node.Val
			count++

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		ans = append(ans, float64(total) / float64(count))
		p = q
	}

	return ans
}