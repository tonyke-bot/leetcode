package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/find-bottom-left-tree-value
Test Case:
[2,1,3]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	leftMost := 0

	layer := []*TreeNode{}
	if root != nil {
		layer = append(layer, root)
	}

	for len(layer) > 0 {
		nextLayer := make([]*TreeNode, 0, len(layer)*2)
		leftMost = layer[0].Val
		for _, node := range layer {
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		layer = nextLayer
	}

	return leftMost
}
