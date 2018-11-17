package leetcode

import (
	"github.com/thagki9/leetcode/kit"
)

// Source: https://leetcode.com/problems/add-two-numbers

// ListNode is a list node
type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	listNode := &ListNode{}
	current := listNode
	carry := 0

	for {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		sum %= 10

		current.Val = sum
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		current.Next = &ListNode{}
		current = current.Next
	}

	return listNode
}

/*
Test Case:
[2,4,3]
[5,6,4]
*/
