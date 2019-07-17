package problem0993

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/cousins-in-binary-tree
Test Case:
[1,2,3,4]
4
3
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	found := 0
	foundDepth := 0
	var foundParent *TreeNode
	matched := false
	target := [101]bool{false}
	target[x] = true
	target[y] = true

	var verify func(node *TreeNode, parentNode *TreeNode, depth int)
	verify = func(node *TreeNode, parentNode *TreeNode, depth int) {
		if node == nil {
			return
		}

		if target[node.Val] {
			// node found
			found++
			if found == 1 {
				foundDepth = depth
				foundParent = parentNode
			} else {
				matched = foundDepth == depth && foundParent != parentNode
			}
			return
		}

		if found == 2 || found == 1 && depth == foundDepth {
			return
		}
		verify(node.Left, node, depth+1)

		if found == 2 || found == 1 && depth == foundDepth {
			return
		}
		verify(node.Right, node, depth+1)
	}

	verify(root, nil, 0)
	return matched

}
