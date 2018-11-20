package leetcode

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_flatten(t *testing.T) {
	root1 := kit.NewTreeNodeFromString("1,2,5,3,4,null,6")
	flatten(root1)
	root1.PrintDFS()

	root2 := kit.NewTreeNodeFromString("2,1,4,null,null,3")
	flatten(root2)
	root2.PrintDFS()
}
