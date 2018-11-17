package leetcode

import "github.com/thagki9/leetcode/kit"

// Source: https://leetcode.com/problems/same-tree

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	same := (p == nil) == (q == nil)
	if !same {
		return false
	}

	if p != nil {
		same = p.Val == q.Val &&
			isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}

	return same
}

/*
Test Case:
[1,2,3]
[1,2,3]
*/
