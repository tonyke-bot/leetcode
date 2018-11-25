package leetcode

import (
	"github.com/thagki9/leetcode/kit"
)

/*
Source: https://leetcode.com/problems/middle-of-the-linked-list
Test Case:
[1,2,3,4,5]

Solution:
	Slow pointer: move one step each iteration
	Fast pointer: move two steps each iteration
*/

// ListNode is a list node
type ListNode = kit.ListNode

func middleNode(head *ListNode) *ListNode {
	slow := head

	for head != nil && head.Next != nil {
		slow = slow.Next
		head = head.Next.Next
	}

	return slow
}
