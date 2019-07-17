package problem0114

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/flatten-binary-tree-to-linked-list
Test Case:
[1,2,5,3,4,null,6]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	originalRight := root.Right
	if root.Left != nil {
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil
		root = moveToEnd(root)
	}

	if originalRight != nil {
		flatten(originalRight)
		root.Right = originalRight
	}
}

func moveToEnd(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	for root.Right != nil {
		root = root.Right
	}

	return root
}
