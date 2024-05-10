type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroupB(head *ListNode, k int) *ListNode {
	// 双指针
	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	end := dummy

	for end.Next != nil {
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				break
			}
		}
		if end == nil {
			break
		}

		// 准备翻转，分组的原始头尾节点
		start := pre.Next
		next := end.Next

		end.Next = nil // 让翻转有结束节点判断
		pre.Next = reverseB(pre.Next)

		// 恢复连接节点
		start.Next = next

		pre = start
		end = start
	}

	return dummy.Next
}

func reverseB(head *ListNode) *ListNode {
	curt := head
	var prev *ListNode = nil
	for curt != nil {
		t := curt.Next
		curt.Next = prev
		prev = curt
		curt = t
	}
	return prev
}
