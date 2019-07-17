package problem1022

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers
Test Case:
[1,0,1,0,1,0,1]
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
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return _sumRootToLeaf(root, 0)
}

func _sumRootToLeaf(root *TreeNode, currentValue int) int {
	currentValue = currentValue<<1 + (root.Val & 1)

	switch {
	case root.Left != nil && root.Right != nil:
		return _sumRootToLeaf(root.Left, currentValue) + _sumRootToLeaf(root.Right, currentValue)
	case root.Left == nil && root.Right == nil:
		return currentValue
	default:
		nextNode := root.Left
		if root.Right != nil {
			nextNode = root.Right
		}
		return _sumRootToLeaf(nextNode, currentValue)
	}
}
