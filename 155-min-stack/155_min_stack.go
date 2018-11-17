package leetcode

// Source: https://leetcode.com/problems/min-stack

// MinStack is
type MinStack struct {
	stack   []int
	minimum []int
	pos     int
}

// Constructor is
func Constructor() MinStack {
	const initialSize = 256
	return MinStack{
		stack:   make([]int, 0, initialSize),
		minimum: make([]int, 0, initialSize),
		pos:     -1,
	}
}

// Push is
func (minStack *MinStack) Push(x int) {
	minStack.stack = append(minStack.stack, x)
	minStack.pos++

	nextMin := x
	if minStack.pos > 0 && minStack.minimum[minStack.pos-1] < x {
		nextMin = minStack.minimum[minStack.pos-1]
	}

	minStack.minimum = append(minStack.minimum, nextMin)
}

// Pop is
func (minStack *MinStack) Pop() {
	if minStack.pos > -1 {
		minStack.minimum = minStack.minimum[:minStack.pos]
		minStack.stack = minStack.stack[:minStack.pos]
		minStack.pos--
	}
}

// Top is
func (minStack *MinStack) Top() int {
	if minStack.pos > -1 {
		return minStack.stack[minStack.pos]
	}
	return 0
}

// GetMin is
func (minStack *MinStack) GetMin() int {
	if minStack.pos > -1 {
		return minStack.minimum[minStack.pos]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
/*
Test Case:
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
*/
