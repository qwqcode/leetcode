func merge(nums1 []int, m int, nums2 []int, n int)  {
	sorted := []int{}

	p1 := 0
	p2 := 0

	for p1 < m || p2 < n {
		if p1 >= m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}

		if p2 >= n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}

	copy(nums1, sorted)
}