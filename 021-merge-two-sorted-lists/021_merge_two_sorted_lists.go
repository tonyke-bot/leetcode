package leetcode

// Source: https://leetcode.com/problems/merge-two-sorted-lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var root *ListNode
	if l1.Val > l2.Val {
		root = l2
		l2 = l2.Next
	} else {
		root = l1
		l1 = l1.Next
	}

	current := root
merge:
	for {
		switch {
		case l1 == nil:
			current.Next = l2
			break merge
		case l2 == nil:
			current.Next = l1
			break merge
		case l1 != nil && l2 != nil:
			if l1.Val > l2.Val {
				current.Next = l2
				l2 = l2.Next
			} else {
				current.Next = l1
				l1 = l1.Next
			}
		default:
			break merge
		}

		current = current.Next
	}

	return root
}

/*
Test Case:
[1,2,4]
[1,3,4]
*/
