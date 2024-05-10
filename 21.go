/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curt := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curt.Next = list1
			list1 = list1.Next
		} else {
			curt.Next = list2
			list2 = list2.Next
		}
		curt = curt.Next
	}
	if list1 != nil {
		curt.Next = list1
	} else {
		curt.Next = list2
	}

	return dummy.Next
}

// 递归法
func mergeTwoListsR(list1 *ListNode, list2 *ListNode) *ListNode {

}
