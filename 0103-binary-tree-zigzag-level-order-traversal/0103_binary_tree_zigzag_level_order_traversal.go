package problem0103

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal
Test Case:
[3,9,20,null,null,15,7]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	values := make([][]int, 0)

	layer := []*TreeNode{}
	if root != nil {
		layer = append(layer, root)
	}
	needReversed := false

	for len(layer) > 0 {
		value := make([]int, 0, len(layer))
		nextLayer := make([]*TreeNode, 0, len(layer)*2)

		for _, node := range layer {
			value = append(value, node.Val)

			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		if needReversed {
			length := len(value)
			for i := 0; i < length-1-i; i++ {
				value[i], value[length-1-i] = value[length-1-i], value[i]
			}
		}

		needReversed = !needReversed
		layer = nextLayer
		values = append(values, value)
	}

	return values
}
