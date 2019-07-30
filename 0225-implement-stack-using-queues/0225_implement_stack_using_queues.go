package problem0225

/*
Source: https://leetcode.com/problems/implement-stack-using-queues
Test Case:
["MyStack","push","push","top","pop","empty"]
[[],[1],[2],[],[],[]]
*/

// MyStack is a stack implemented with queues.
type MyStack struct {
	queue []int
}

// Constructor initializes your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

// Push element x onto stack. */
func (s *MyStack) Push(x int) {
	s.queue = append(s.queue, x)
}

// Pop removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	tempQ := make([]int, 0, len(s.queue)-1)
	for len(s.queue) > 1 {
		tempQ = append(tempQ, s.queue[0])
		s.queue = s.queue[1:]
	}

	val := s.queue[0]
	s.queue = tempQ
	return val
}

// Top gets the top element. */
func (s *MyStack) Top() int {
	val := s.Pop()
	s.Push(val)
	return val
}

// Empty returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
