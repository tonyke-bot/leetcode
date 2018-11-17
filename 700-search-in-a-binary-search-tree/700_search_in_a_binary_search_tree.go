package leetcode

import (
	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/search-in-a-binary-search-tree
Test Case:
[4,2,7,1,3]
2
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}

		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}
