package problem0965

import (
	"fmt"
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_isUnivalTree(t *testing.T) {
	tests := []struct {
		args *TreeNode
		want bool
	}{
		{kit.NewTreeNodeFromString("9,9,6,9,9"), false},
		{kit.NewTreeNodeFromString("1,1,1,1,1,null,1"), true},
		{kit.NewTreeNodeFromString("2,2,2,5,2"), false},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case #%d", index), func(t *testing.T) {
			if got := isUnivalTree(tt.args); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
