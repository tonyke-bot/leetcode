package problem0009

// Source: https://leetcode.com/problems/palindrome-number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == original
}

/*
Test Case:
121
*/
