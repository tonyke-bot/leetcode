package main

// Source: https://leetcode.com/problems/remove-element

func removeElement(nums []int, val int) int {
	for pos := 0; pos < len(nums); pos++ {
		if nums[pos] != val {
			continue
		}

		copy(nums[pos:], nums[pos+1:])
		nums = nums[:len(nums)-1]
		pos--
	}

	return len(nums)
}

/*
Test Case:
[3,2,2,3]
3
*/
