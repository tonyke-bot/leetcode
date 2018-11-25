package leetcode

import (
	"testing"

	"github.com/thagki9/leetcode/kit"
)

func Test_splitListToParts(t *testing.T) {
	splitListToParts(kit.NewListNode([]int{1, 2, 3, 4}), 5)
}
