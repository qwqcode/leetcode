func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	
	ans := 0
	n := len(s)
	for i := range s {
		v := m[s[i]]

		if i < n-1 && v < m[s[i+1]] {
			ans -= v
		} else {
			ans += v
		}
	}

	return ans
}