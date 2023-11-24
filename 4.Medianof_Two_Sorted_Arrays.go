package leetcode

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sorted := make([]int, 0, len(nums1)+len(nums2))
	sorted = append(sorted, nums1...)
	sorted = append(sorted, nums2...)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return findMedian(sorted)
}

func findMedian(nums []int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	if len(nums) == 2 {
		return float64(nums[0]+nums[1]) / 2
	}
	return findMedian(nums[1 : len(nums)-1])
}
