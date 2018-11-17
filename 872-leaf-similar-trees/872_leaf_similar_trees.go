package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/leaf-similar-trees
Test Case:
[3,5,1,6,2,9,8,null,null,7,4]
[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := []int{}
	getLeaf(root1, &leaves1)
	return compareLeaves(root2, &leaves1) && len(leaves1) == 0
}

func compareLeaves(root *TreeNode, leavesToCompare *[]int) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		if len(*leavesToCompare) == 0 || (*leavesToCompare)[0] != root.Val {
			return false
		}

		*leavesToCompare = (*leavesToCompare)[1:]
		return true
	}

	return compareLeaves(root.Left, leavesToCompare) && compareLeaves(root.Right, leavesToCompare)
}

func getLeaf(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	} else {
		getLeaf(root.Left, leaves)
		getLeaf(root.Right, leaves)
	}
}
