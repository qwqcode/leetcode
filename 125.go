func isLetter(i rune) bool {
	return ((i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z'))
}

func isPalindrome(s string) bool {
	w := ""
	for _, e := range s {
		if isLetter(e) {
			w += strings.ToLower(string(e))
		}
	}

	fmt.Println(len(w))

	r := len(w) - 1
	l := 0

	for l < r {
		fmt.Println(l, r, w[l], w[r])
		if w[l] != w[r] {
			return false
		}

		l++
		r--
	}

	return true
}
