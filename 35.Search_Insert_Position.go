package leetcode

func searchInsert(nums []int, target int) int {
	low := 0
	hight := len(nums) - 1
	middle := low + (hight-low)/2

	for low < hight {

		if target > nums[middle] {
			low = middle + 1
		} else if target < nums[middle] {
			hight = middle - 1

		} else {
			return middle
		}
		middle = low + (hight-low)/2
	}

	if target <= nums[middle] && middle == 0 {
		return 0
	} else if (target < nums[middle] && middle != 0) || target == nums[middle] {
		return middle
	} else {
		return middle + 1
	}

}
