package problem0671

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_findSecondMinimumValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case #01", args{kit.NewTreeNodeFromString("1,1,3,1,1,3,4,3,1,1,1,3,8,4,8,3,3,1,6,2,1")}, 2},
		{"case #02", args{kit.NewTreeNodeFromString("2,2,5,null,null,5,7")}, 5},
		{"case #03", args{kit.NewTreeNodeFromString("2,2,2")}, -1},
		{"case #04", args{kit.NewTreeNodeFromString("2,2,5,2,3,5,7")}, 3},
		{"case #05", args{kit.NewTreeNodeFromString("2,2,3,2,5,3,7")}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
