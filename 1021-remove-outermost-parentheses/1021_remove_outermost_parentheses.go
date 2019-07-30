package problem1021

/*
Source: https://leetcode.com/problems/remove-outermost-parentheses
Test Case:
"(()())(())"
*/

func removeOuterParentheses(S string) string {
	stackSize := 0
	result := ""
	sequence := ""

	chars := []byte(S)
	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case '(':
			stackSize++
		case ')':
			stackSize--
		}
		sequence += string(chars[i])

		if stackSize == 0 {
			result += sequence[1 : len(sequence)-1]
			sequence = ""
		}
	}

	return result
}
