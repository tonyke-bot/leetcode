package problem0561

import "sort"

/*
Source: https://leetcode.com/problems/array-partition-i
Test Case:
[1,4,3,2]
*/

type nums561 []int

func (nums nums561) Len() int           { return len(nums) }
func (nums nums561) Swap(i, j int)      { nums[i], nums[j] = nums[j], nums[i] }
func (nums nums561) Less(i, j int) bool { return nums[i] < nums[j] }

func arrayPairSum(nums []int) int {
	sort.Sort(nums561(nums))

	sum := 0
	for i := len(nums) - 2; i >= 0; i -= 2 {
		sum += nums[i]
	}

	return sum
}
