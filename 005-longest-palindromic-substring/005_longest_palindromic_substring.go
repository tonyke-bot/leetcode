package main

// Source: https://leetcode.com/problems/longest-palindromic-substring
/*
Solution:
	Double pointer
*/

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	result := ""
	maxLength := 0

	for left := 0; left < length-1; left++ {
		right := length - 1

		for right >= left && right-left+1 > maxLength {
			if s[right] == s[left] {
				tempLeft, tempRight := left+1, right-1
				isPalidrome := true

				for isPalidrome && tempLeft < tempRight {
					isPalidrome = s[tempRight] == s[tempLeft]
					tempRight--
					tempLeft++
				}

				if isPalidrome {
					result = s[left : right+1]
					maxLength = right - left + 1
					break
				}
			}

			right--
		}
	}

	return result
}

/*
Test Case:
"babad"
*/
