package problem0199

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/binary-tree-right-side-view
Test Case:
[1,2,3,null,5,null,4]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func rightSideView(root *TreeNode) []int {
	values := make([]int, 0)

	layer := []*TreeNode{}
	if root != nil {
		layer = append(layer, root)
	}

	for len(layer) > 0 {
		nextLayer := make([]*TreeNode, 0, len(layer)*2)
		rightMost := layer[0].Val

		for _, node := range layer {
			rightMost = node.Val
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		values = append(values, rightMost)
		layer = nextLayer
	}

	return values
}
