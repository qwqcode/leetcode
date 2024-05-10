func isHappy(n int) bool {
	fast := n
	slow := n

	for {
		fast = calcSum(fast)
		fast = calcSum(fast)
		slow = calcSum(slow)

		if fast == slow {
			break
		}
	}
	
	return slow == 1
}

func calcSum(n int) int {
	sum := 0
	for n > 0 {
		bit := n % 10
		sum += bit * bit
		n = n / 10
	}
	return sum
}
