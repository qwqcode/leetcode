func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}


	t := 0
	for t < x {
		t = t * 10 + x % 10
		x = x / 10
	}

	return t == x || t / 10 == x
}