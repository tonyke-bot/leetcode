package main

// Source: https://leetcode.com/problems/valid-parentheses

var pair = map[byte]byte{'}': '{', ']': '[', ')': '('}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for _, char := range []byte(s) {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pair[char] {
				return false
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

/*
Test Case:
"()"
*/
