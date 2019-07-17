package problem0725

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/split-linked-list-in-parts
Test Case:
[1,2,3,4]
5
*/

// ListNode is a list node
type ListNode = kit.ListNode

func splitListToParts(root *ListNode, k int) []*ListNode {
	if k == 0 {
		return nil
	} else if k == 1 {
		return []*ListNode{root}
	}

	tmpPointer := root
	length := 0
	for tmpPointer != nil {
		tmpPointer = tmpPointer.Next
		length++
	}

	sectorLength := length / k
	specialSector := length % k

	result := make([]*ListNode, 0, k)
	for k > 0 {
		head := &ListNode{}
		tmpPointer = head
		count := sectorLength
		if specialSector > 0 {
			count++
			specialSector--
		}

		for count > 0 {
			tmpPointer.Next = root
			tmpPointer = tmpPointer.Next
			root = tmpPointer.Next
			count--
		}

		tmpPointer.Next = nil
		result = append(result, head.Next)
		k--
	}

	return result
}
