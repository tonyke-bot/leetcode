package leetcode

/*
Source: https://leetcode.com/problems/longest-continuous-increasing-subsequence
Test Case:
[1,3,5,4,7]
*/

func findLengthOfLCIS(nums []int) int {
	maxLength := 0
	currentLength := 0
	lastNumber := -1 << 31

	for _, num := range nums {
		if num > lastNumber {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}

		lastNumber = num
	}

	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}
