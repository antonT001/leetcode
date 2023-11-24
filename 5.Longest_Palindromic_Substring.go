package leetcode

func longestPalindrome(s string) string {
	var (
		startPointer = 0
		endPointer   = len(s) - 1
		result       string
	)

	if len(s) == 1 {
		return s
	}

	for startPointer < len(s) {
		if s[endPointer] != s[startPointer] {
			endPointer--
		} else if startPointer == endPointer {
			if len(result) < 1 {
				result = string(s[startPointer])
			}
			startPointer++
			endPointer = len(s) - 1
		} else if s[endPointer] == s[startPointer] {
			find := s[startPointer : endPointer+1]
			if isPalindrome(find) {
				if len(result) < len(find) {
					result = find
				}
				startPointer++
				endPointer = len(s) - 1
			} else {
				endPointer--
			}
		}
	}

	return result
}

func isPalindrome(s string) bool {
	var (
		startPointer = 0
		endPointer   = len(s) - 1
	)

	for startPointer < endPointer {
		if s[startPointer] != s[endPointer] {
			return false
		}

		startPointer++
		endPointer--
	}

	return true
}
