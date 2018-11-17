package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/average-of-levels-in-binary-tree
Test Case:
[3,9,20,15,7]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	layer := []*TreeNode{}
	result := []float64{}

	if root != nil {
		layer = append(layer, root)
	}

	for len(layer) > 0 {
		nextLayer := make([]*TreeNode, 0, len(layer)*2)
		sum := int64(0)

		for _, node := range layer {
			sum += int64(node.Val)
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		result = append(result, float64(sum)/float64(len(layer)))
		layer = nextLayer
	}

	return result
}
