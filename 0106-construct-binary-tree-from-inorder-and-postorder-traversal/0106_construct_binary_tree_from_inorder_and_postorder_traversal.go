package problem0106

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
Test Case:
[9,3,15,20,7]
[9,15,7,20,3]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) != len(inorder) {
		return nil
	}

	length := len(postorder)
	if length == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[length-1]}

	if length > 1 {
		mid := 0
		for inorder[mid] != root.Val {
			mid++
		}

		root.Left = buildTree(inorder[:mid], postorder[:mid])
		root.Right = buildTree(inorder[mid+1:], postorder[mid:length-1])
	}

	return root
}
