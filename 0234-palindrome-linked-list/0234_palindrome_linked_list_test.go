package problem0234

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{kit.NewListNode([]int{1, 2, 3, 2, 1})}, true},
		{"", args{kit.NewListNode([]int{1, 2, 3, 3, 2, 1})}, true},
		{"", args{kit.NewListNode([]int{1, 2, 4, 3, 2, 1})}, false},
		{"", args{kit.NewListNode([]int{1, 2, 2, 1})}, true},
		{"", args{kit.NewListNode([]int{1, 1})}, true},
		{"", args{kit.NewListNode([]int{1})}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
