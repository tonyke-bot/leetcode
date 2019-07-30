package problem1047

/*
Source: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string
Test Case:
"abbaca"
*/

func removeDuplicates(S string) string {
	bytes := []byte(S)
	stack := [20001]byte{} // byte at index 0 is abondoned to avoid extra check in loop
	stackPos := 0

	for i := 0; i < len(bytes); i++ {
		if stack[stackPos] == bytes[i] {
			stackPos--
		} else {
			stackPos++
			stack[stackPos] = bytes[i]
		}
	}

	return string(stack[1 : stackPos+1])
}
