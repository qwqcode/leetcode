func candy(ra []int) int {
	need := make([]int, len(ra))

	for i := 0; i < len(ra); i++ {
		need[i] = 1
	}

	for i := 0; i < len(ra) -1; i++ {
		if ra[i+1] > ra[i] {
			need[i+1] = need[i] + 1
		}
	}

	for i := len(ra) - 1; i > 0; i-- {
		if ra[i-1] > ra[i] {
			need[i-1] = max(need[i] + 1, need[i-1])
		}
	}

	sum := 0
	for i := 0; i < len(need); i++ {
		sum += need[i]
	}

	return sum
}