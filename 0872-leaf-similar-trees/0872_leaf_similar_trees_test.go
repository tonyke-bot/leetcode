package problem0872

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{kit.NewTreeNodeFromString("3,5,1,6,2,9,8,null,null,7,4"), kit.NewTreeNodeFromString("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
