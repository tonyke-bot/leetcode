package problem0889

import (
	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal
Test Case:
[1,2,4,5,3,6,7]
[4,5,2,6,7,3,1]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) != len(post) || len(pre) == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}
	length := len(pre)

	if length > 1 {
		leftRoot := length - 2
		for post[leftRoot] != pre[1] {
			leftRoot--
		}

		root.Left = constructFromPrePost(pre[1:leftRoot+2], post[:leftRoot+1])
		root.Right = constructFromPrePost(pre[leftRoot+2:], post[leftRoot+1:length-1])
	}

	return root
}
