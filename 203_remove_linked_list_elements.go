package leetcode

/*
Source: https://leetcode.com/problems/remove-linked-list-elements
Test Case:
[1,2,6,3,4,5,6]
6
*/

func removeElements(head *ListNode, val int) *ListNode {
	virtualHead := &ListNode{Next: head}
	head = virtualHead

	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return virtualHead.Next
}
