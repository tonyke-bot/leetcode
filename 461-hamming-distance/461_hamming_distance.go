package main

// Source: https://leetcode.com/problems/hamming-distance

func hammingDistance(x int, y int) int {
	count := 0
	tempX := uint(x)
	tempY := uint(y)

	for pit := uint(1); pit <= 1<<32; pit <<= 1 {
		if tempX&pit != tempY&pit {
			count++
		}
	}

	return count
}

/*
Test Case:
1
4
*/
