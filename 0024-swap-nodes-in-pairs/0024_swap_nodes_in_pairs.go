package problem0024

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/swap-nodes-in-pairs
Test Case:
[1,2,3,4]
*/

// ListNode is a list node
type ListNode = kit.ListNode

func swapPairs(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	pointer := fakeHead

	for pointer.Next != nil && pointer.Next.Next != nil {
		node1 := pointer.Next
		node2 := node1.Next
		node3 := node2.Next

		pointer.Next = node2
		pointer.Next.Next = node1
		pointer.Next.Next.Next = node3

		pointer = node1
	}

	return fakeHead.Next
}
