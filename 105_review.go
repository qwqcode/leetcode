/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	index := map[int]int{}
	for pos, num := range inorder {
		index[num] = pos
	}

	var myBuildTree func(preOrderLeft, preOrderRight, inOrderLeft, inOrderRight int) *TreeNode
	myBuildTree = func(preOrderLeft, preOrderRight, inOrderLeft, inOrderRight int) *TreeNode {
		if preOrderLeft > preOrderRight {
			return nil
		}

		preOrderRoot := preOrderLeft
		inOrderRoot := index[preorder[preOrderRoot]]

		sizeLeftSubtree := inOrderRoot - inOrderLeft // 左子树的长度

		root := &TreeNode{inorder[inOrderRoot], nil, nil}
		root.Left = myBuildTree(preOrderLeft + 1, preOrderLeft + sizeLeftSubtree, inOrderLeft, inOrderRoot - 1)
		root.Right = myBuildTree(preOrderLeft + sizeLeftSubtree + 1, preOrderRight, inOrderRoot + 1, inOrderRight)

		return root
	}

	return myBuildTree(0, n - 1, 0, n - 1)
}