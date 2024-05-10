/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
        return nil
    }
    
	dummy := ListNode{Val: -1}
	dummy.Next = head

	curt := dummy

	for curt.Next != nil && curt.Next.Next != nil {
		val := curt.Next.Val

		if curt.Next.Next.Val == val {
			for curt.Next != nil && curt.Next.Val == val {
				curt.Next = curt.Next.Next
			}
		} else {
			curt = curt.Next
		}
	}

	return dummy.Next
}