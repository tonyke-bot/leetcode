package leetcode

// Source: https://leetcode.com/problems/two-sum

func twoSum(nums []int, target int) []int {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		find := target - nums[i]
		for j := i + 1; j < size; j++ {
			if nums[j] == find {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
Test Case:
[2,7,11,15]
9
*/
