func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])

		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func lcp(s1, s2 string) string {
	n := min(len(s1), len(s2))

	ans := ""
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			ans += string(s1[i])
		} else {
			break
		}
	}

	return ans
}