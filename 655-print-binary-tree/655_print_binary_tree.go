package leetcode

import (
	"strconv"

	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/print-binary-tree
Test Case:
[1,2]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	width := 1<<uint(height) - 1

	if height == 0 {
		return nil
	}

	result := make([][]string, height)
	layer := []*TreeNode{root}

	for depth := 0; depth < height; depth++ {
		nextLayer := make([]*TreeNode, 0, len(layer)*2)
		row := make([]string, 0, width)

		prefixSize := (1 << uint(height-depth-1)) - 1
		gapSize := (1 << uint(height-depth)) - 1

		for i := 0; i < prefixSize; i++ {
			row = append(row, "")
		}

		nodeCount := len(layer)
		for i := 0; i < nodeCount; i++ {
			node := layer[i]

			if node == nil {
				row = append(row, "")
				nextLayer = append(nextLayer, nil, nil)
			} else {
				row = append(row, strconv.Itoa(node.Val))
				nextLayer = append(nextLayer, node.Left, node.Right)
			}

			if i < nodeCount-1 {
				for j := 0; j < gapSize; j++ {
					row = append(row, "")
				}
			}
		}

		for i := 0; i < prefixSize; i++ {
			row = append(row, "")
		}

		result[depth] = row
		layer = nextLayer
	}

	return result
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
