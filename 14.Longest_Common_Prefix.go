package leetcode

func longestCommonPrefix(strs []string) string {
	var (
		result  string
		element byte
	)

	for j := 0; ; j++ {
		for i := range strs {
			if j > len(strs[i])-1 {
				return result
			}

			if element == 0 {
				element = strs[i][j]
			} else if element != strs[i][j] {
				return result
			}

		}
		result += string(element)
		element = 0
	}
}
