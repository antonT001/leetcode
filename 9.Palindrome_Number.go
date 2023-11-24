package leetcode

func isPalindrome2(x int) bool {
	newX := 0

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	for x > newX {
		newX *= 10
		newX += x % 10
		x /= 10
	}

	if x == newX || x == newX/10 {
		return true
	}
	return false
}
