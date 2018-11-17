package leetcode

/*
Source: https://leetcode.com/problems/power-of-four
Test Case:
16

Solution:
	No recursion, No loop. Should consider bit operation

	Observe power of 4 in binary

	4^0   0b1
	4^1   0b100
	4^2   0b10000
	4^3   0b1000000

	Power of 4 in binary has two property and verification:
	1. Only one 1
		4^1 - 1       = 0b011
		4^1           = 0b100
		(4^1-1)&(4^1) = 0b000

	2. Even amount of trailing zero
		4^1 - 1       = 0b011
		4^2 - 1       = 0b01111
		4^n - 1       = 0b11 << (n-1)*2 _ 0b11 = 3 * 4(n-1) + 3, is the multiples of 3

		if there is odd amount of 1, for example
			0b1 = 1
			0b111 = 7
			0b11111 = 31
			...
			0b111 << 2 + 0b111
		Obviously is not the multiples of 3

	Hence, if we can verify these two properties of a number, we can tell that this
		number is the power of 4

*/

func isPowerOfFour(num int) bool {
	return num > 0 && (num-1)&num == 0 && (num-1)%3 == 0
}
