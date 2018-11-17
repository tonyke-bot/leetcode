package main

/*
Source: https://leetcode.com/problems/move-zeroes
Test Case:
[0,1,0,3,12]
*/

func moveZeroes(nums []int) {
	insertPos := 0

	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}

	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}

/* Brute Force
func moveZeroes(nums []int) {
	zeroPos := len(nums) - 1
	for zeroPos >= 0 && nums[zeroPos] == 0 {
		zeroPos--
	}

	if zeroPos < 0 {
		return
	}

	pos := 0
	for pos <= zeroPos {
		if nums[pos] == 0 {
			for start := pos + 1; start <= zeroPos; start++ {
				nums[start-1] = nums[start]
			}
			nums[zeroPos] = 0
			zeroPos--
		} else {
			pos++
		}
	}
}
*/
