package leetcode

// Source: https://leetcode.com/problems/minimum-depth-of-binary-tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	switch {
	case root.Left == nil && root.Right == nil:
		return 1
	case root.Left == nil:
		return minDepth(root.Right) + 1
	case root.Right == nil:
		return minDepth(root.Left) + 1
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	switch {
	case leftDepth < rightDepth:
		return leftDepth + 1
	default:
		return rightDepth + 1
	}
}

/*
Test Case:
[3,9,20,null,null,15,7]
*/
