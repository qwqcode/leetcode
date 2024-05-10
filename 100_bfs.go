func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	q1, q2 := []*TreeNode{p}, []*TreeNode{q}

	for len(q1) > 0 && len(q2) > 0 {
		n1, n2 := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:]
		if n1.Val != n2.Val {
			return false
		}
		l1, r1 := n1.Left, n1.Right
		l2, r2 := n2.Left, n2.Right
		if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
			return false
		}
		if (r1 == nil && r2 != nil) || (r1 != nil && r2 == nil) {
			return false
		}
		if l1 != nil {
			q1 = append(q1, l1)
		}
		if r1 != nil {
			q1 = append(q1, r1)
		}
		if l2 != nil {
			q2 = append(q2, l2)
		}
		if r2 != nil {
			q2 = append(q2, r2)
		}
	}

	return len(q1) == 0 && len(q2) == 0
}