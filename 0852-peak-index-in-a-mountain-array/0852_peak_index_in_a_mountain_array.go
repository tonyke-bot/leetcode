package problem0852

/*
Source: https://leetcode.com/problems/peak-index-in-a-mountain-array
Test Case:
[0,1,0]
*/

func peakIndexInMountainArray(A []int) int {
	length := len(A)
	for i := 0; i < length-1; i++ {
		if A[i] > A[i+1] {
			return i
		}
	}
	return length - 1
}
