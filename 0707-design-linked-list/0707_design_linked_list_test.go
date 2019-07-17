package problem0707

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	fmt.Println(list.Get(1))
	list.DeleteAtIndex(1)
	fmt.Println(list.Get(1))
}
