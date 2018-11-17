package leetcode

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_tree2str(t *testing.T) {
	type args struct {
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{kit.NewTreeNodeFromString("1,2,3,4")}, "1(2(4))(3)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.t); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
