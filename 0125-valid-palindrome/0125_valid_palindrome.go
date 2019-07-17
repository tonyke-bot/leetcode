package problem0125

// Source: https://leetcode.com/problems/valid-palindrome

func isPalindromeII(s string) bool {
	bytes := []byte(s)
	left := 0
	right := len(bytes) - 1

	for left < right {
		// skip left side non-alphanumberic character
		for left < right && !(('a' <= bytes[left] && bytes[left] <= 'z') || ('A' <= bytes[left] && bytes[left] <= 'Z') || ('0' <= bytes[left] && bytes[left] <= '9')) {
			left++
		}
		// skip right side non-alphanumberic character
		for left < right && !(('a' <= bytes[right] && bytes[right] <= 'z') || ('A' <= bytes[right] && bytes[right] <= 'Z') || ('0' <= bytes[right] && bytes[right] <= '9')) {
			right--
		}

		if left < right {
			if bytes[left]|0x20 != bytes[right]|0x20 {
				return false
			}
			left++
			right--
		}
	}

	return true
}

/*
Test Case:
"A man, a plan, a canal: Panama"
*/
