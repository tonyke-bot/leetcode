package leetcode

/*
Source: https://leetcode.com/problems/sort-array-by-parity-ii
Test Case:
[4,2,5,7]
*/

func sortArrayByParityII(A []int) []int {
	result := make([]int, len(A))

	even := 0
	odd := 1

	for _, v := range A {
		if v%2 == 0 {
			result[even] = v
			even += 2
		} else {
			result[odd] = v
			odd += 2
		}
	}

	return result
}
