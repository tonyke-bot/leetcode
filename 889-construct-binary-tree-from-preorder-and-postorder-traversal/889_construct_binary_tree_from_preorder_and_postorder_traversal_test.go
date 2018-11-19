package leetcode

import (
	"testing"
)

var root3 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 6},
		Right: &TreeNode{Val: 7},
	},
}

func Test_constructFromPrePost(t *testing.T) {
	constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}).PrintDFS()
}
