package problem0058

import (
	"strings"
)

// Source: https://leetcode.com/problems/length-of-last-word

func lengthOfLastWord(s string) int {
	bytes := []byte(strings.Trim(s, " "))

	for pos := len(bytes) - 1; pos >= 0; pos-- {
		switch {
		case ('A' <= bytes[pos] && bytes[pos] <= 'Z') || ('a' <= bytes[pos] && bytes[pos] <= 'z'):
			continue
		case bytes[pos] == ' ':
			return len(bytes) - pos - 1
		}
	}

	return len(bytes)
}

/*
Test Case:
"Hello World"
*/
