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
    for i, num := range inorder {
        index[num] = i
    }

    var myBuildTree func(preOrderLeft, preOrderRight, inOrderLeft, inOrderRight int) *TreeNode
    myBuildTree = func (preOrderLeft, preOrderRight, inOrderLeft, inOrderRight int) *TreeNode {
        if preOrderLeft > preOrderRight {
            return nil
        }

        // 前序遍历的第一个节点就是根节点
        preOrderRoot := preOrderLeft

        // 从索引中找到中序遍历中的根节点下标
        inOrderRoot := index[preorder[preOrderRoot]]

        // 得到左子树的节点数目（在中序遍历中的）
        sizeLeftSubtree := inOrderRoot - inOrderLeft

        // 建立根节点
        root := &TreeNode{preorder[preOrderRoot], nil, nil}

        root.Left = myBuildTree(preOrderLeft + 1, preOrderLeft + sizeLeftSubtree, inOrderLeft, inOrderRoot - 1)
        root.Right = myBuildTree(preOrderLeft + sizeLeftSubtree + 1, preOrderRight, inOrderRoot + 1, inOrderRight)
        return root
    }


    return myBuildTree(0, n - 1, 0, n - 1)
}