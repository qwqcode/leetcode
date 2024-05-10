func reverseList(head *ListNode) {
	curt := head
	var prev *ListNode

	for curt != nil {
		t := curt.Next
		curt.Next = prev
		prev = curt
		curt = t
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 因为头节点有变化的可能，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 截取
	succ := rightNode.Next
	leftNode := pre.Next
	pre.Next = nil
	rightNode.Next = nil

	// 翻转链表子区间
	reverseList(leftNode)

	// 接回去
	pre.Next = rightNode
	leftNode.Next = succ

	return dummyNode.Next
}