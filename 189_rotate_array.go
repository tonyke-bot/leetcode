package leetcode

/*
Source: https://leetcode.com/problems/rotate-array
Solution:
	1. Divide and concat: O(n) time, O(n) space
	2. Simulate step by step: O(n*k) time, O(1) space
	3. Copy subarrays to another array: O(n) time, O(n) space
*/

func rotate(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}
	k %= length
	if k == 0 {
		return
	}

	tmp := make([]int, 0, length)
	tmp = append(tmp, nums[length-k:]...)
	tmp = append(tmp, nums[:length-k]...)
	copy(nums, tmp)
}

/*
Test Case:
[1,2,3,4,5,6,7]
3
*/
