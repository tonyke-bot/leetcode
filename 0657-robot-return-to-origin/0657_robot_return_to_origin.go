package problem0657

// Source: https://leetcode.com/problems/robot-return-to-origin

func judgeCircle(moves string) bool {
	horizontal := 0
	vertical := 0

	for _, direction := range []byte(moves) {
		switch direction {
		case 'U':
			vertical++
		case 'D':
			vertical--
		case 'R':
			horizontal++
		case 'L':
			horizontal--
		}
	}

	return horizontal == 0 && vertical == 0
}

/*
Test Case:
"UD"
*/
