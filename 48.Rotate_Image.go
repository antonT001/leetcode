package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		for k := 0; k < n/2; k++ {
			matrix[i][k], matrix[i][n-1-k] = matrix[i][n-1-k], matrix[i][k]
		}
	}
}
