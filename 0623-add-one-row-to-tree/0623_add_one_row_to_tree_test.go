package problem0623

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_addOneRow(t *testing.T) {
	// addOneRow(kit.NewTreeNodeFromString("1,2,3"), 10, 1).PrintDFS()
	// addOneRow(kit.NewTreeNodeFromString("4,2,null,3,1"), 1, 3).PrintDFS()
	// addOneRow(kit.NewTreeNodeFromString("1,2,3,4"), 1, 4).PrintDFS()
	addOneRow(kit.NewTreeNodeFromString("4,2,6,3,1,5"), 1, 3).PrintDFS()
}
