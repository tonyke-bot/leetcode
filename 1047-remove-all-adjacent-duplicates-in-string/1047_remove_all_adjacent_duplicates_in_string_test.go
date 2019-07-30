package problem1047

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"aaaaaaaaa", "a"},
		{"aababaab", "ba"},
		{"abbaca", "ca"},
		{"abbacca", "a"},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case %d", index), func(t *testing.T) {
			if got := removeDuplicates(tt.args); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
