package leetcode

func romanToInt(s string) int {
	var (
		romanMap = map[byte]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}

		result  int
		lastNum int
	)

	for i := len(s) - 1; i >= 0; i-- {
		num := romanMap[s[i]]
		if num >= lastNum {
			result += num
		} else {
			result -= num
		}
		lastNum = num
	}
	return result
}
