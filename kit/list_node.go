package kit

import "fmt"

// ListNode is a link node
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode creates a list node from given array
func NewListNode(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	l := &ListNode{}
	current := l
	for _, v := range values {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	return l.Next
}

// Print prints the link list, from start to end.
// Output ends with comma.
func (l *ListNode) Print() {
	if l == nil {
		fmt.Println(nil)
		return
	}

	for ; l != nil; l = l.Next {
		fmt.Print(l.Val)
		fmt.Print(", ")
	}
	fmt.Print("\n")
}
