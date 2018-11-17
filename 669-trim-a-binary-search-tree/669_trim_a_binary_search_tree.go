package leetcode

import (
	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/trim-a-binary-search-tree
Test Case:
[1,0,2]
1
2
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < L {
		return trimBST(root.Right, L, R)
	} else if root.Val > R {
		return trimBST(root.Left, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
