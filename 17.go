var phoneMap map[string]string = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

    ans := []string{}
	path := []byte{}

	var dfs func(depth int)
	dfs = func(depth int) {
		if depth == len(digits) {
			ans = append(ans, string(path))
			return
		}

		digit := string(digits[depth])
		letters := phoneMap[digit]
		
		for i := range letters {
			path = append(path, letters[i])
			dfs(depth + 1)
			path = path[:len(path) - 1]
		}
	}

	dfs(0)

	return ans
}