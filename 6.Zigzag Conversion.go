package leetcode

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var buffer bytes.Buffer
	buffer.Grow(len(s))

	step := numRows*2 - 2
	for i := 0; i < numRows; i++ {
		subStep := step - i*2
		for j := i; j < len(s); j += step {
			buffer.WriteByte(s[j])
			if i != 0 && i != numRows-1 && j+subStep < len(s) {
				buffer.WriteByte(s[j+subStep])
			}
		}
	}
	return buffer.String()
}
