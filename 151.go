func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	ans := ""
	
	for _, w := range arr {
		if w == "" {
			continue
		}

		ans = w + " " + ans
	}

	return strings.TrimSpace(ans)
}
