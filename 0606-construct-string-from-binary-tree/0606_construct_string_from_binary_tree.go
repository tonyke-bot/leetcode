package problem0606

import (
	"strconv"

	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/construct-string-from-binary-tree
Test Case:
[1,2,3,4]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func tree2str(t *TreeNode) string {
	result := ""

	if t != nil {
		result = strconv.Itoa(t.Val)
		leftResult := tree2str(t.Left)
		rightResult := tree2str(t.Right)

		if rightResult != "" {
			result += "(" + leftResult + ")(" + rightResult + ")"
		} else if leftResult != "" {
			result += "(" + leftResult + ")"
		}
	}

	return result
}
