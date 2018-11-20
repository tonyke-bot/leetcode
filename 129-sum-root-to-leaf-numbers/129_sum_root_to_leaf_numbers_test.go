package leetcode

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{kit.NewTreeNodeFromString("1,2,3")}, 25},
		{"", args{kit.NewTreeNodeFromString("1,0")}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
