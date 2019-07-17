package problem0771

// Source: https://leetcode.com/problems/jewels-and-stones

func numJewelsInStones(J string, S string) int {
	isJewels := new([256]bool)
	for _, v := range []byte(J) {
		isJewels[v] = true
	}

	count := 0
	for _, v := range []byte(S) {
		if isJewels[v] {
			count++
		}
	}
	return count
}

/*
Test Case:
"aA"
"aAAbbbb"
*/
