func strStr(haystack string, needle string) int {
	for i := 0; i + len(needle) <= len(haystack); i++ {
		f := true

		for j := 0; j < len(needle); j++ {
			if haystack[i + j] != needle[j] {
				f = false
				break
			}
		}

		if f {
			return i
		}
	}

	return -1
}
