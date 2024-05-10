func canCompleteCircuit(gas []int, cost []int) int {
	i := 0
	n := len(gas)

	for i < n {
		sumGas := 0
		sumCost := 0
		t := 0
		
		for t < n {
			j := (i + t) % n
			sumGas += gas[j]
			sumCost += cost[j]

			if sumCost > sumGas {
				break
			}
			t++
		}

		if t == n {
			return i
		} else {
			i = i + t + 1
		}
	}

	return -1
}