package problem0965

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/univalued-binary-tree
Test Case:
[1,1,1,1,1,null,1]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	return _isUnivalTree(root, nil)
}

func _isUnivalTree(node *TreeNode, parent *TreeNode) bool {
	if node == nil {
		return true
	}

	if parent != nil {
		if node.Val != parent.Val {
			return false
		}
	}

	if node.Left != nil && !_isUnivalTree(node.Left, node) {
		return false
	}

	if node.Right != nil && !_isUnivalTree(node.Right, node) {
		return false
	}

	return true
}
