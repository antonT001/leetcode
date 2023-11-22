package leetcode

func lengthOfLongestSubstring(s string) int {
	set := [128]bool{}
	lenght := 0
	res := 0
	l := 0
	for _, v := range s {
		for ; set[v]; l++ {
			lenght--
			set[s[l]] = false
		}
		set[v] = true
		lenght++
		if res < lenght {
			res = lenght
		}
	}
	return res
}
