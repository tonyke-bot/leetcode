package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/palindrome-linked-list
Test Case:
[1,2]

Solution:
	I implemented a space O(1) solution by reverse the first half
	of the linked list, during the comparison I recover the original
	order of that half of list.
*/

// ListNode is a list node
type ListNode = kit.ListNode

func isPalindrome(head *ListNode) bool {
	switch {
	case head == nil || head.Next == nil:
		return true
	case head.Next != nil && head.Next == nil:
		return head.Val == head.Next.Val
	}

	var tmpPointer *ListNode

	// Check if it's odd length
	isOddLength := false
	tmpPointer = head
	for tmpPointer != nil {
		tmpPointer = tmpPointer.Next
		isOddLength = !isOddLength
	}

	// Find middle of the link node
	mid := head
	tmpPointer = head
	for tmpPointer != nil && tmpPointer.Next != nil {
		mid = mid.Next
		tmpPointer = tmpPointer.Next.Next
	}

	// Reverse first half of link node
	var lastNode *ListNode
	for head != mid {
		head.Next, lastNode, head = lastNode, head, head.Next
	}

	allSame := true

	// Compare and reverse
	list1 := lastNode
	list2 := mid
	if isOddLength {
		list2 = list2.Next
	}

	lastNode = mid
	for list2 != nil {
		allSame = allSame && (list1.Val == list2.Val)
		list1, lastNode, list1.Next = list1.Next, list1, lastNode
		list2 = list2.Next
	}

	return allSame
}
