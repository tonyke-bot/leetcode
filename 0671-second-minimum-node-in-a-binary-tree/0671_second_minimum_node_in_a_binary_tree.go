package problem0671

import (
	"math"

	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/second-minimum-node-in-a-binary-tree
Test Case:
[2,2,5,null,null,5,7]
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
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}

	const pesudoMin2 = int64(math.MaxInt64)
	min2 := pesudoMin2
	dfs(root, root.Val, &min2)

	if min2 == pesudoMin2 {
		return -1
	}
	return int(min2)
}

func dfs(node *TreeNode, min int, min2 *int64) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}

	if node.Left.Val == node.Right.Val {
		dfs(node.Left, min, min2)
		dfs(node.Right, min, min2)
		return
	}

	sameValueNode := node.Left
	differentValueNode := node.Right
	if node.Right.Val == min {
		sameValueNode = node.Right
		differentValueNode = node.Left
	}

	if int64(differentValueNode.Val) < *min2 {
		*min2 = int64(differentValueNode.Val)
		dfs(sameValueNode, min, min2)
	}
}
