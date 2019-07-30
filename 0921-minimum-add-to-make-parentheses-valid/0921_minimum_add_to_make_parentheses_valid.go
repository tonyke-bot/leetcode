package problem0921

/*
Source: https://leetcode.com/problems/minimum-add-to-make-parentheses-valid
Test Case:
"())"
*/

func minAddToMakeValid(S string) int {
	balance := 0
	neededParentheses := 0

	for _, char := range []byte(S) {
		if char == '(' {
			if balance < 0 {
				neededParentheses += -balance
				balance = 0
			}
			balance++
		} else {
			balance--
		}
	}

	if balance < 0 {
		balance = -balance
	}
	neededParentheses += balance
	return neededParentheses
}
