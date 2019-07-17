package problem0538

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/convert-bst-to-greater-tree
Test Case:
[5,2,13]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	convertBSTWithWeight(root, 0)
	return root
}

func convertBSTWithWeight(root *TreeNode, weight int) int {
	sum := 0

	if root != nil {
		sum += root.Val
		root.Val += weight

		rightSum := convertBSTWithWeight(root.Right, weight)
		root.Val += rightSum
		sum += rightSum

		leftSum := convertBSTWithWeight(root.Left, sum+weight)
		sum += leftSum
	}

	return sum
}
