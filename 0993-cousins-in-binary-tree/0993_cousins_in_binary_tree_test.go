package problem0993

import (
	"fmt"
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_isCousins(t *testing.T) {
	tests := []struct {
		root *TreeNode
		x    int
		y    int
		want bool
	}{
		{kit.NewTreeNodeFromString("1,2,3,4"), 4, 3, false},
		{kit.NewTreeNodeFromString("1,2,3,null,4,null,5"), 5, 4, true},
		{kit.NewTreeNodeFromString("1,2,3,null,4"), 2, 3, false},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case #%d", index), func(t *testing.T) {
			if got := isCousins(tt.root, tt.x, tt.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
