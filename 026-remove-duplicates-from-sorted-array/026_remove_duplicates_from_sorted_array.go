package main

// Source: https://leetcode.com/problems/remove-duplicates-from-sorted-array

func removeDuplicates(nums []int) int {
	lastNumber := 1<<32 - 1
	lastPos := -1

	for pos := 0; pos < len(nums); pos++ {
		num := nums[pos]
		if lastNumber == num {
			continue
		}

		if pos != 0 {
			copy(nums[lastPos+1:], nums[pos:])
			nums = nums[:len(nums)-(pos-lastPos-1)]
		}

		pos = lastPos + 1
		lastNumber = num
		lastPos++
	}

	if len(nums) > 0 && nums[len(nums)-1] == lastNumber && len(nums)-1 > lastPos {
		nums = nums[:lastPos+1]
	}

	return len(nums)
}

/*
Test Case:
[1,1,2]
*/
