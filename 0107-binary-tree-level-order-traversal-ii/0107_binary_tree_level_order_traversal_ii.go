package problem0107

import "github.com/thagki9/leetcode/kit"

// Source: https://leetcode.com/problems/binary-tree-level-order-traversal-ii

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	pyramid := make([][]int, 0, 1024)
	layer := []*TreeNode{root}

	for {
		nums := make([]int, 0, len(layer))
		nextLayer := make([]*TreeNode, 0, len(layer)*2)

		for _, node := range layer {
			if node != nil {
				nextLayer = append(nextLayer, node.Left)
				nextLayer = append(nextLayer, node.Right)
				nums = append(nums, node.Val)
			}
		}

		if len(nums) == 0 {
			break
		}

		layer = nextLayer
		pyramid = append(pyramid, nums)
	}

	for i := 0; i < len(pyramid)-1-i; i++ {
		pyramid[i], pyramid[len(pyramid)-1-i] = pyramid[len(pyramid)-1-i], pyramid[i]
	}

	return pyramid
}

/*
Test Case:
[3,9,20,null,null,15,7]
*/
