package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/invert-binary-tree
Test Case:
[4,2,7,1,3,6,9]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}

	return root
}
