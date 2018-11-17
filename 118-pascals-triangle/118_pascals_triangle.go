package leetcode

// Source: https://leetcode.com/problems/pascals-triangle

func generate(numRows int) [][]int {
	pyramid := [][]int{
		[]int{1},
		[]int{1, 1},
	}

	switch numRows {
	case 0:
		return nil
	case 1, 2:
		return pyramid[0:numRows]
	}

	for i := 2; i < numRows; i++ {
		lastLayer := pyramid[i-1]
		layer := append(make([]int, 0, len(lastLayer)+1), 1)

		for j := 1; j < i; j++ {
			layer = append(layer, lastLayer[j-1]+lastLayer[j])
		}

		layer = append(layer, 1)
		pyramid = append(pyramid, layer)
	}

	return pyramid
}

/*
Test Case:
5
*/
