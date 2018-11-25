package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/linked-list-components
Test Case:
[0,1,2,3]
[0,1,3]
*/

// ListNode is a list node
type ListNode = kit.ListNode

func numComponents(head *ListNode, G []int) int {
	set := [10000]bool{}
	for _, item := range G {
		set[item] = true
	}

	count := 0
	current := 0

	for head != nil {
		if set[head.Val] {
			if current == 0 {
				count++
			}
			current++
		} else {
			current = 0
		}
		head = head.Next
	}

	return count
}
