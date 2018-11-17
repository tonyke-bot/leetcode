package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/binary-tree-tilt
Test Case:
[1,2,3]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func findTilt(root *TreeNode) int {
	tilt, _ := _findTitl(root)
	return tilt
}

func _findTitl(root *TreeNode) (tilt, sum int) {
	if root != nil {
		tiltLeft, sumLeft := _findTitl(root.Left)
		tiltRight, sumRight := _findTitl(root.Right)

		tilt = tiltLeft + tiltRight + abs(sumLeft-sumRight)
		sum = sumLeft + sumRight + root.Val
	}
	return
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
