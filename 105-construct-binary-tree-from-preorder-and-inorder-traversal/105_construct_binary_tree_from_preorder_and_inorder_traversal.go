package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
Test Case:
[3,9,20,15,7]
[9,3,15,20,7]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	length := len(preorder)
	if length == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	if length > 1 {
		mid := 0
		for inorder[mid] != root.Val {
			mid++
		}

		root.Left = buildTree(preorder[1:1+mid], inorder[:mid])
		root.Right = buildTree(preorder[1+mid:], inorder[mid+1:])
	}

	return root
}
