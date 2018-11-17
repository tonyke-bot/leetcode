package main

/*
Source: https://leetcode.com/problems/valid-palindrome-ii
Test Case:
"aba"
*/

func valid(bytes []byte, left, right int) bool {
	for left < right {
		if bytes[left] != bytes[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func validPalindrome(s string) bool {
	bytes := []byte(s)
	left := 0
	right := len(bytes) - 1

	for left < right {
		if bytes[left] != bytes[right] {
			return valid(bytes, left+1, right) || valid(bytes, left, right-1)
		}

		left++
		right--
	}

	return true
}
