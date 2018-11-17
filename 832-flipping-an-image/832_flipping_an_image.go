package leetcode

// Source: https://leetcode.com/problems/flipping-an-image

func flipAndInvertImage(A [][]int) [][]int {
	image := make([][]int, 0, len(A))

	for _, row := range A {
		newRow := make([]int, 0, len(row))

		for i := len(row) - 1; i >= 0; i-- {
			newRow = append(newRow, (row[i]+1)%2)
		}

		image = append(image, newRow)
	}

	return image
}

/*
Test Case:
[[1,1,0],[1,0,1],[0,0,0]]
*/
