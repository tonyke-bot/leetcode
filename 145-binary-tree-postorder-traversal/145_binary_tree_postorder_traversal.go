package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/binary-tree-postorder-traversal
Test Case:
[1,null,2,3]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root != nil {
		dfs(root.Left, result)
		dfs(root.Right, result)
		*result = append(*result, root.Val)
	}
}
