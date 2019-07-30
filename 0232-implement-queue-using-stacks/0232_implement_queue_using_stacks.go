package problem0232

/*
Source: https://leetcode.com/problems/implement-queue-using-stacks
Test Case:
["MyQueue","push","push","peek","pop","empty"]
[[],[1],[2],[],[],[]]
*/

// MyQueue is a queue implemented with stacks.
type MyQueue struct {
	stackIn  []int
	stackOut []int
}

// Constructor initializes your data structure here.
func Constructor() MyQueue {
	return MyQueue{}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

// Pop removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	val := q.Peek()
	if len(q.stackOut) > 0 {
		q.stackOut = q.stackOut[:len(q.stackOut)-1]
	}
	return val
}

// Peek gets the front element.
func (q *MyQueue) Peek() int {
	q.moveFromInToOut()
	if len(q.stackOut) == 0 {
		return 0
	}
	return q.stackOut[len(q.stackOut)-1]
}

// Empty returns whether the queue is empty.
func (q *MyQueue) Empty() bool {
	return len(q.stackIn) == 0 && len(q.stackOut) == 0
}

func (q *MyQueue) moveFromInToOut() {
	if len(q.stackIn) == 0 || len(q.stackOut) > 0 {
		return
	}

	for i := len(q.stackIn) - 1; i >= 0; i-- {
		q.stackOut = append(q.stackOut, q.stackIn[i])
	}

	q.stackIn = nil
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
