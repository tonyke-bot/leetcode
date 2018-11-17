package leetcode

import "github.com/thagki9/leetcode/kit"

// Source: https://leetcode.com/problems/balanced-binary-tree

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func isBalanced(root *TreeNode) bool {
	delta := calcDepthDelta(root)
	return delta != -1
}

func calcDepthDelta(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := calcDepthDelta(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := calcDepthDelta(root.Right)
	if rightDepth == -1 {
		return -1
	}
	delta := leftDepth - rightDepth

	if !(-1 <= delta && delta <= 1) {
		return -1
	}

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

/*
Test Case:
[3,9,20,null,null,15,7]
*/
