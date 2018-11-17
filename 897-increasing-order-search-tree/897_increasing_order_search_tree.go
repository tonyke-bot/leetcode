package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/increasing-order-search-tree
Test Case:
[5,3,6,2,4,null,8,1,null,null,null,7,9]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	result := &TreeNode{}
	result.Right = increasingBST(root.Left)

	rightMost := result
	for rightMost.Right != nil {
		rightMost = rightMost.Right
	}

	rightMost.Right = &TreeNode{
		Val:   root.Val,
		Right: increasingBST(root.Right),
	}

	return result.Right
}
