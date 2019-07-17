package problem0129

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/sum-root-to-leaf-numbers
Test Case:
[1,2,3]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, 0)
}

func dfs(root *TreeNode, prefix int) int {
	prefix = prefix*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return prefix
	}

	sum := 0
	if root.Left != nil {
		sum += dfs(root.Left, prefix)
	}
	if root.Right != nil {
		sum += dfs(root.Right, prefix)
	}
	return sum
}
