/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	x := head
	for x.Next != nil {
		x = x.Next
		n++
	}

	x.Next = head

	times := n - k % n

	for times != 0 {
		x = x.Next
		times--
	}

	ret := x.Next
	x.Next = nil

	return ret
}
