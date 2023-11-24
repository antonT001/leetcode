package leetcode

import (
	"slices"
)

func maxCoins(piles []int) int {
	slices.Sort(piles)

	step := len(piles) / 3
	result := 0
	for i := len(piles)-2; step != 0; i -= 2 {
		result += piles[i]
		step--
	}
	return result
}
