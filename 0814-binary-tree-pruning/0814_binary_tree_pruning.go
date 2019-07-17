package problem0814

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/binary-tree-pruning
Test Case:
[1,null,0,0,1]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
