package leetcode

// Source: https://leetcode.com/problems/sort-array-by-parity

func sortArrayByParity(A []int) []int {
	odd := make([]int, 0, len(A))
	even := make([]int, 0, len(A))
	for _, v := range A {
		if v%2 == 1 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
		}
	}
	return append(even, odd...)
}

/*
Test Case:
[3,1,2,4]
*/
