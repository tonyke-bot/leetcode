package problem0108

import "github.com/thagki9/leetcode/kit"

// Source: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree
// Solution:
//   1. Choose root(mid point)
//   2. Construct left tree
//   3. Construct right tree
//   4. Do Step2 & Step3 recursively
//
//   In my first submission, I use only one function to do DFS with the
//     help of slice. But I only beat one digit percent of Go solutions.
//   And I doubt that slice has bad performance so I add another function
//     which has parameters 'left' and 'right' to indicate the array boundry.
//   The below solution beat 98% of Go solutions.

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, left, right int) *TreeNode {
	mid := (left + right) >> 1
	root := &TreeNode{Val: nums[mid]}
	if mid > left {
		root.Left = dfs(nums, left, mid-1)
	}
	if mid+1 <= right {
		root.Right = dfs(nums, mid+1, right)
	}

	return root
}

/*
Test Case:
[-10,-3,0,5,9]
*/
