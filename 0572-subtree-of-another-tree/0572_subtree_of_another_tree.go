package problem0572

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/subtree-of-another-tree
Test Case:
[3,4,5,1,2]
[4,1,2]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return s != nil &&
		t != nil &&
		((s.Val == t.Val && isSameTree(s, t)) || isSubtree(s.Left, t) || isSubtree(s.Right, t))
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	return (s == nil) == (t == nil) &&
		(s == nil || (s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)))
}
