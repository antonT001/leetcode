package leetcode

import "math"

func reverse(x int) int {
	var (
		isNegative bool
		result     uint32
	)

	if x < 0 {
		isNegative = true
		x *= -1
	}

	for x > 0 {
		tmp := uint32(x % 10)
		x /= 10
		a := result
		result *= 10
		result += tmp

		if a != (result-tmp)/10 || result > math.MaxInt32 {
			return 0
		}
	}

	if isNegative {
		return int(result) * -1
	}

	return int(result)
}
