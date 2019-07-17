package problem0654

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/maximum-binary-tree
Test Case:
[3,2,1,6,0,5]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := 0
	for i, num := range nums {
		if num > nums[maxIndex] {
			maxIndex = i
		}
	}

	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}
