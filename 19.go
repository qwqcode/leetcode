func getLength(head *ListNode) int {
	i := 0
	curt := head

	for curt != nil {
		i++
		curt = curt.Next
	}

	return i
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	curt := dummy

	for i := 0; i < length-n; i++ {
		curt = curt.Next
	}
	curt.Next = curt.Next.Next
	return dummy.Next
}