package main

// Source: https://leetcode.com/problems/longest-common-prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	pos := 0
	chars := make([]byte, 0, 1024)
	firstLength := len(strs[0])

	for {
		if pos >= firstLength {
			break
		}

		char := strs[0][pos]
		same := true

		for _, str := range strs[1:] {
			same = pos < len(str) && str[pos] == char
			if !same {
				break
			}
		}

		if !same {
			break
		}

		pos++
		chars = append(chars, char)
	}

	return string(chars)
}

/*
Test Case:
["flower","flow","flight"]
*/
