package problem0102

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/binary-tree-level-order-traversal
Test Case:
[3,9,20,null,null,15,7]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func levelOrder(root *TreeNode) [][]int {
	values := make([][]int, 0)

	layer := []*TreeNode{}
	if root != nil {
		layer = append(layer, root)
	}

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

		layer = nextLayer
		values = append(values, value)
	}

	return values
}
