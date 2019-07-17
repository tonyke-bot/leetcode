package problem0572

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{kit.NewTreeNodeFromString("1,null,1,null,1,null,1,null,1,null,1,2"), kit.NewTreeNodeFromString("1,null,1,null,1,null,1,null,1,null,1,2")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{kit.NewTreeNodeFromString("1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,2"), kit.NewTreeNodeFromString("1,null,1,null,1,null,1,null,1,null,1,2")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
