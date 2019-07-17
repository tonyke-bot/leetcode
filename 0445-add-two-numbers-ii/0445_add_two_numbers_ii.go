package problem0445

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/add-two-numbers-ii
Test Case:
[7,2,4,3]
[5,6,4]
*/

// ListNode is a list node
type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	stack2 := []int{}

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	carry := 0
	var lastNode *ListNode
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		sum := carry
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		lastNode = &ListNode{Val: sum % 10, Next: lastNode}
		carry = sum / 10
	}

	return lastNode
}
