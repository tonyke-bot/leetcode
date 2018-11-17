package main

// Source: https://leetcode.com/problems/plus-one

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	buffer := make([]int, 0, len(digits)+1)
	buffer = append(buffer, digits[len(digits)-1]+1)

	carry := buffer[0] / 10
	buffer[0] %= 10

	for i := len(digits) - 2; i >= 0; i-- {
		num := digits[i] + carry
		buffer = append(buffer, num%10)
		carry = num / 10
	}

	if carry > 0 {
		buffer = append(buffer, carry)
	}

	for i := 0; i < len(buffer)-1-i; i++ {
		buffer[i], buffer[len(buffer)-1-i] = buffer[len(buffer)-1-i], buffer[i]
	}

	return buffer
}

/*
Test Case:
[1,2,3]
*/
