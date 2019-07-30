package problem0225

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(1)
	if popVal := s.Pop(); popVal != 1 {
		t.Errorf("popVal = %d, expected = %d", popVal, 1)
	}
	if popVal := s.Pop(); popVal != 1 {
		t.Errorf("popVal = %d, expected = %d", popVal, 1)
	}
	if isEmpty := s.Empty(); !isEmpty {
		t.Errorf("s.Empty() = %v, expected = %v", isEmpty, true)
	}
}
