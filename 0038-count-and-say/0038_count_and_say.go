package problem0038

// Source: https://leetcode.com/problems/count-and-say

func countAndSay(n int) string {
	say := []byte("1")

	for ; n > 1; n-- {
		buffer := make([]byte, 0, len(say)*2)

		lastChar := say[0]
		count := 1

		for pos := 1; pos < len(say); pos++ {
			if say[pos] == lastChar {
				count++
			} else {
				buffer = append(buffer, byte(count+'0'))
				buffer = append(buffer, lastChar)
				lastChar = say[pos]
				count = 1
			}
		}

		buffer = append(buffer, byte(count+'0'))
		buffer = append(buffer, lastChar)
		say = buffer
	}

	return string(say)
}

/*
Test Case:
1
*/
