package leetcode

import "sort"

func twoSum(nums []int, target int) []int {

	numsCp := make([]int, len(nums))
	copy(numsCp, nums)
	sort.Ints(numsCp)
	for i := 0; i < len(numsCp); i++ {
		j := sort.SearchInts(numsCp, target-numsCp[i])

		if j < len(numsCp) && numsCp[j] == target-numsCp[i] && i != j {
			numsCp[0] = numsCp[i]
			numsCp[1] = numsCp[j]
			break
		}

	}
	for i := 0; i < len(nums); i++ {
		if numsCp[0] == nums[i] {
			numsCp[0] = i
			break
		}
	}
	for i := 0; i < len(nums); i++ {
		if numsCp[1] == nums[i] && i != numsCp[0] {
			numsCp[1] = i
			break
		}
	}
	return numsCp[:2]
}
