func isAnagram(s string, t string) bool {
    m1 := map[byte]int{}
    m2 := map[byte]int{}
	if len(s) != len(t) {
		return false
	}
	n := len(s)
	for i := 0; i < n; i++ {
		m1[s[i]]++
		m2[t[i]]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}