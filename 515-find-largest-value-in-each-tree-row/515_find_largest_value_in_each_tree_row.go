package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/find-largest-value-in-each-tree-row
Test Case:
[1,3,2,5,3,null,9]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func largestValues(root *TreeNode) []int {
	values := make([]int, 0)

	layer := []*TreeNode{}
	if root != nil {
		layer = append(layer, root)
	}

	for len(layer) > 0 {
		maxValue := -1 << 31
		nextLayer := make([]*TreeNode, 0, len(layer)*2)

		for _, node := range layer {
			if node.Val > maxValue {
				maxValue = node.Val
			}

			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		layer = nextLayer
		values = append(values, maxValue)
	}

	return values
}
