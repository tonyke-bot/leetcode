package problem0112

import "github.com/thagki9/leetcode/kit"

// Source: https://leetcode.com/problems/path-sum

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Right == nil && root.Left == nil {
		return sum == root.Val
	}

	sum -= root.Val
	return (root.Left != nil && hasPathSum(root.Left, sum)) ||
		(root.Right != nil && hasPathSum(root.Right, sum))
}

/*
Test Case:
[5,4,8,11,null,13,4,7,2,null,null,null,1]
22
*/
