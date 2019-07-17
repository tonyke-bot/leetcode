package problem0113

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/path-sum-ii
Test Case:
[5,4,8,11,null,13,4,7,2,null,null,5,1]
22
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	paths := [][]int{}
	if root != nil {
		walk(root, sum, make([]int, 0, 32), &paths)
	}
	return paths
}

func walk(root *TreeNode, sum int, path []int, paths *[][]int) {
	sum -= root.Val
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			*paths = append(*paths, newPath)
		}
		return
	}

	if root.Left != nil {
		walk(root.Left, sum, path, paths)
	}

	if root.Right != nil {
		walk(root.Right, sum, path, paths)
	}
}
