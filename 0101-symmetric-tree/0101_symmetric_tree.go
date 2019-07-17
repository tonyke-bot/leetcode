package problem0101

import "github.com/thagki9/leetcode/kit"

// Source: https://leetcode.com/problems/symmetric-tree

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		return ((left == nil) == (right == nil)) &&
			(left == nil ||
				(left.Val == right.Val &&
					dfs(left.Left, right.Right) &&
					dfs(left.Right, right.Left)))
	}

	return dfs(root.Left, root.Right)
}

/*
Test Case:
[1,2,2,3,4,4,3]
*/
