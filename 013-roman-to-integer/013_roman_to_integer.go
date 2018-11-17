package leetcode

// Source: https://leetcode.com/problems/roman-to-integer

var values = map[int]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	chars := []byte(s)
	pos := len(chars) - 1
	result := 0

	// I can be placed before V (5) and X (10) to make 4 and 9.
	// X can be placed before L (50) and C (100) to make 40 and 90.
	// C can be placed before D (500) and M (1000) to make 400 and 900.

	for pos >= 0 {
		char := chars[pos]
		result += values[int(char)]
		switch char {
		case 'V', 'X':
			if pos-1 >= 0 && chars[pos-1] == 'I' {
				result--
				pos--
			}
		case 'L', 'C':
			if pos-1 >= 0 && chars[pos-1] == 'X' {
				result -= 10
				pos--
			}
		case 'D', 'M':
			if pos-1 >= 0 && chars[pos-1] == 'C' {
				result -= 100
				pos--
			}
		}
		pos--
	}

	return result
}

/*
Test Case:
"III"
*/
