package leetcode

// Source: https://leetcode.com/problems/climbing-stairs
// Solution:
//   This problem can be solved using DP
//   D[n] stands for how many ways to climb to n
//   D[n] = | D[n-1] + D[n-2], n > 2
//          | 2, n = 2
//          | 1, n = 1
//          | 0, n = 0
func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	lastSecond := 0
	lastFirst := 1

	for ; n > 1; n-- {
		lastSecond, lastFirst = lastFirst, lastFirst+lastSecond
	}

	return lastFirst
}

/*
Test Case:
2
*/
