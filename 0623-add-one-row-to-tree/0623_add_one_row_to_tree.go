package problem0623

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/add-one-row-to-tree
Test Case:
[4,2,6,3,1,5]
1
2
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d < 1 {
		return nil
	}

	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}

	layer := []*TreeNode{}
	if root != nil {
		layer = append(layer, root)
	}

	for ; d > 2; d-- {
		nextLayer := make([]*TreeNode, 0, len(layer)*2)
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

	isLastLayer := true
	for _, node := range layer {
		isLastLayer = node.Left == nil || node.Right == nil
		if !isLastLayer {
			break
		}
	}

	for _, node := range layer {
		node.Left = &TreeNode{Val: v, Left: node.Left}
		node.Right = &TreeNode{Val: v, Right: node.Right}
	}

	return root
}
