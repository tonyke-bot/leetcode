package problem0257

import (
	"strconv"

	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/binary-tree-paths
Test Case:
[1,2,3,null,5]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	dfs(root, "", &result)
	return result
}

func dfs(root *TreeNode, path string, paths *[]string) {
	if root == nil {
		return
	}

	if path != "" {
		path += "->"
	}
	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
		return
	}

	dfs(root.Left, path, paths)
	dfs(root.Right, path, paths)
}
