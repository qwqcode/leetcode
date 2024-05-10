/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{Val: -1}
	smallH := small // 哨兵节点

	large := &ListNode{Val: -1}
	largeH := large

	iter := head
	for iter != nil {
		if iter.Val < x {
			small.Next = iter
			small = small.Next
		} else {
			large.Next = iter
			large = large.Next
		}

		iter = iter.Next
	}

	large.Next = nil
	small.Next = largeH.Next

	return smallH.Next
}