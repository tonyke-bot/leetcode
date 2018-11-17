package leetcode

/*
Source: https://leetcode.com/problems/toeplitz-matrix
Test Case:
[[1,2,3,4],[5,1,2,3],[9,5,1,2]]
*/

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for y := m - 2; y > 0; y-- {
		left := matrix[y][0]

		for x := 1; y+x < m && x < n; x++ {
			if matrix[y+x][x] != left {
				return false
			}
		}
	}

	for x := 0; x < n-1; x++ {
		pivot := matrix[0][x]
		for y := 1; y < m && x+y < n; y++ {
			if pivot != matrix[y][x+y] {
				return false
			}
		}
	}

	return true
}
