package main

import (
	"strings"
)

/*
Source: https://leetcode.com/problems/reverse-words-in-a-string-iii
Test Case:
"Let's take LeetCode contest"
*/

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	result := make([]string, 0, len(words))

	for _, word := range words {
		buffer := make([]byte, len(word))
		bytes := []byte(word)
		for index, char := range bytes {
			buffer[len(bytes)-1-index] = char
		}
		result = append(result, string(buffer))
	}

	return strings.Join(result, " ")
}
