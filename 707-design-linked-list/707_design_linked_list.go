package leetcode

/*
Source: https://leetcode.com/problems/design-linked-list
Test Case:
["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
[[],[1],[3],[1,2],[1],[1],[1]]
*/

type node struct {
	next *node
	val  int
}

// MyLinkedList is my listed list
type MyLinkedList struct {
	head *node
	tail *node
	size int
}

// Constructor initializes your data structure here.
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (list *MyLinkedList) Get(index int) int {
	if index >= list.size {
		return -1
	}

	head := list.head
	for index > 0 {
		head = head.next
		index--
	}

	return head.val
}

// AddAtHead adds a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (list *MyLinkedList) AddAtHead(val int) {
	list.head = &node{val: val, next: list.head}
	if list.size == 0 {
		list.tail = list.head
	}
	list.size++
}

// AddAtTail appends a node of value val to the last element of the linked list.
func (list *MyLinkedList) AddAtTail(val int) {
	newNode := &node{val: val}
	if list.size == 0 {
		list.tail, list.head = newNode, newNode
	} else {
		list.tail, list.tail.next = newNode, newNode
	}

	list.size++
}

// AddAtIndex adds a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index > list.size:
		return
	case index == list.size:
		list.AddAtTail(val)
		return
	case index == 0:
		list.AddAtHead(val)
		return
	}

	head := list.head
	for index > 1 {
		head = head.next
		index--
	}

	head.next = &node{val: val, next: head.next}
	list.size++
}

// DeleteAtIndex deletes the index-th node in the linked list, if the index is valid.
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if list.size == 0 || index >= list.size {
		return
	}

	switch {
	case list.size == 1:
		list.head = nil
		list.tail = nil
	case index == 0:
		list.head.next = nil
		list.head = list.head.next
	default:
		head := list.head
		for index > 1 {
			head = head.next
			index--
		}

		head.next, head.next.next = head.next.next, nil
		if head.next == nil {
			list.tail = head
		}
	}

	list.size--
}
