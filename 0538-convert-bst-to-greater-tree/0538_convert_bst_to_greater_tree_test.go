package problem0538

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_convertBST(t *testing.T) {
	convertBST(kit.NewTreeNodeFromString("6,4,9,3,5,8,10,2")).PrintDFS()
}
