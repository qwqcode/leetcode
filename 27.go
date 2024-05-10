func removeElement(nums []int, val int) int {
	idx := 0

	for _, v := range nums {
		if v == val {
			continue
		}

		nums[idx] = v
		idx++
	}

	return idx
}