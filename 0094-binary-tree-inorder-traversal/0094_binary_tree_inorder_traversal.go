package problem0094

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/binary-tree-inorder-traversal
Test Case:
[1,null,2,3]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root != nil {
		dfs(root.Left, result)
		*result = append(*result, root.Val)
		dfs(root.Right, result)
	}
}
