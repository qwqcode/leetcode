func convert(s string, numRows int) string {
	// 记得当行数小于 2 时，直接输出 s
	if numRows < 2 {
		return s
	}

	res := make([]string, numRows)

	i := 0

	// 状态 flag，1 或 -1，当 1 时 i++，当 -1 时 i--
	flag := -1

	for _, b := range s {
		c := string(b)

		res[i] += c

		if i == 0 || i == numRows-1 {
			flag = -flag
		}

		i = i + flag
	}

	return strings.Join(res, "")
}