package leetcode

func removeElement(nums []int, val int) int {
	var (
		j = len(nums) - 1
		i int
	)

	for i <= j {
		if nums[i] == val {
			if nums[j] != val {
				nums[i] = nums[j]
			} else {
				i--
			}
			j--
		}
		i++
	}
	return i
}
