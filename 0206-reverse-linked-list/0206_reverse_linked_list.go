package problem0206

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/reverse-linked-list
Test Case:
[1,2,3,4,5]
*/

// ListNode is a list node
type ListNode = kit.ListNode

func reverseList(head *ListNode) *ListNode {
	var lastNode *ListNode

	for head != nil {
		head.Next, lastNode, head = lastNode, head, head.Next
	}

	return lastNode
}

func reverseListDFS(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lastNode := reverseListDFS(head.Next)
	head.Next.Next = head
	head.Next = nil
	return lastNode
}
