func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := dummy
	end := dummy

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		start := pre.Next

		next := end.Next
		end.Next = nil

		pre.Next = reverse(start)

		start.Next = next

		pre = start
		end = pre
	}

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
