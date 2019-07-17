package problem0894

import (
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
	for _, tree := range allPossibleFBT(5) {
		tree.PrintDFS()
	}
}
