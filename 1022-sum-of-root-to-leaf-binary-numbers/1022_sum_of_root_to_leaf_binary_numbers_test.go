package problem1022

import (
	"fmt"
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_sumRootToLeaf(t *testing.T) {
	tests := []struct {
		args *TreeNode
		want int
	}{
		{kit.NewTreeNodeFromString("1,0,1,0,1,0,1"), 22},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case %d", index), func(t *testing.T) {
			if got := sumRootToLeaf(tt.args); got != tt.want {
				t.Errorf("sumRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
