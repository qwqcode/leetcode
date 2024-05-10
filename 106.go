/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	m := map[int]int{}
	for pos, num := range inorder {
		m[num] = pos
	}
	
	var myBuildTree func(inLeft, inRight, postLeft, postRight int) *TreeNode
	myBuildTree = func(inLeft, inRight, postLeft, postRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}

		postRoot := postRight // 注意这里是 postRight 不是 postLeft
		inRoot := m[postorder[postRoot]]
		size := inRoot - inLeft

		node := &TreeNode{inorder[inRoot], nil, nil}
		node.Left = myBuildTree(inLeft, inRoot - 1, postLeft, postLeft + size - 1)
		node.Right = myBuildTree(inRoot + 1, inRight, postLeft + size, postRight - 1)
		
		return node
	}

	return myBuildTree(0, n - 1, 0, n - 1)
}