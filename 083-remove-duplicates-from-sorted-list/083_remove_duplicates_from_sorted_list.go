package main

// Source: https://leetcode.com/problems/remove-duplicates-from-sorted-list

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for current.Next != nil {
		if current.Next.Val != current.Val {
			current = current.Next
		} else {
			current.Next = current.Next.Next
		}
	}

	return head
}

/*
Test Case:
[1,1,2]
*/
