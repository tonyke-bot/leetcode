package main

import "strconv"

/*
Source: https://leetcode.com/problems/self-dividing-numbers
Test Case:
1
22
*/

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)

	for left <= right {
		isSelfDividing := true
		bytes := []byte(strconv.Itoa(left))
		length := len(bytes)
		for i := 0; i < length && isSelfDividing; i++ {
			num := int(bytes[i]) - '0'
			isSelfDividing = num > 0 && left%num == 0
		}

		if isSelfDividing {
			result = append(result, left)
		}
		left++
	}

	return result
}
