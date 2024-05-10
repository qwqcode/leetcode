/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	queue := []*Node{root}
	
	for len(queue) > 0 {
		nodes := []*Node{}
		
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node == nil {
				continue
			}
			
			if i - 1 >= 0 {
				queue[i-1].Next = node
			}
			
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}

		queue = nodes
	}

	return root
}