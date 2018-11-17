package main

// Source: https://leetcode.com/problems/maximum-subarray

func maxSubArray(nums []int) int {
	maxSum := -1 << 31
	current := 0

	for _, num := range nums {
		current += num

		if current > maxSum {
			maxSum = current
		}

		if current < 0 {
			current = 0
		}
	}

	return maxSum
}

/*
Test Case:
[-2,1,-3,4,-1,2,1,-5,4]
*/
