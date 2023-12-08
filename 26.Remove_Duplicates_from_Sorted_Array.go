package leetcode

func removeDuplicates(nums []int) int {
    var i int
    total := len(nums)
    for j:= 1; j < len(nums); j++ {
        if nums[i] == nums[j] {
            total--
        } else {
            i++
            nums[i]=nums[j]
        }
    }
    return total
}
