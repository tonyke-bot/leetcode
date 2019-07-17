package problem0404

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/sum-of-left-leaves
Test Case:
[3,9,20,null,null,15,7]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	if root != nil {
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				sum += root.Left.Val
			} else {
				sum += sumOfLeftLeaves(root.Left)
			}
		}

		sum += sumOfLeftLeaves(root.Right)
	}

	return sum
}
