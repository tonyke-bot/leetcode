package problem0007

import (
	"math"
)

// Source: https://leetcode.com/problems/reverse-integer

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	sign := x > 0
	reversed := 0

	if !sign {
		x = -x
	}

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	if !sign {
		reversed = -reversed
	}
	return reversed
}

/*
Test Case:
123
*/
