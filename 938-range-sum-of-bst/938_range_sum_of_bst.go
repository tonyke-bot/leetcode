package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/range-sum-of-bst
Test Case:
[10,5,15,3,7,null,18]
7
15
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if R < root.Val {
		return rangeSumBST(root.Left, L, R)
	} else if L > root.Val {
		return rangeSumBST(root.Right, L, R)
	} else {
		return rangeSumBST(root.Left, L, R) + root.Val + rangeSumBST(root.Right, L, R)
	}
}
