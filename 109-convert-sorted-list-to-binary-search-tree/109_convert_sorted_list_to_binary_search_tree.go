package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree
Test Case:
[-10,-3,0,5,9]
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

// ListNode is a list node
type ListNode = kit.ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	var lastMid *ListNode
	mid := head
	tmp := head
	for tmp != nil && tmp.Next != nil {
		lastMid = mid
		mid = mid.Next
		tmp = tmp.Next.Next
	}

	lastMid.Next = nil
	root := &TreeNode{
		Val:   mid.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(mid.Next),
	}
	lastMid.Next = mid

	return root
}
