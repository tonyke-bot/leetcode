package problem0496

/*
Source: https://leetcode.com/problems/next-greater-element-i
Test Case:
[4,1,2]
[1,3,4,2]
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	stack := make([]int, 0, len(nums2))
	nextGreater := make(map[int]int)

	for _, ele := range nums2 {
		for len(stack) > 0 && ele > stack[len(stack)-1] {
			nextGreater[stack[len(stack)-1]] = ele
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ele)
	}

	for i, ele := range nums1 {
		if value, ok := nextGreater[ele]; ok {
			result[i] = value
		} else {
			result[i] = -1
		}
	}

	return result
}
