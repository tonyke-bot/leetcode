package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/all-possible-full-binary-trees
Test Case:
7
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func allPossibleFBT(n int) []*TreeNode {
	completeBTrees := map[int][]*TreeNode{
		1: []*TreeNode{&TreeNode{Val: 0}},
	}

	for node := 3; node <= n; node++ {
		trees := []*TreeNode{}
		for rootPos := 1; rootPos < node; rootPos++ {
			leftN := rootPos - 1
			rightN := node - rootPos

			leftBTrees := completeBTrees[leftN]
			rightBTrees := completeBTrees[rightN]

			if len(leftBTrees) == 0 || len(rightBTrees) == 0 {
				continue
			} else {
				for _, left := range leftBTrees {
					for _, right := range rightBTrees {
						trees = append(trees, &TreeNode{
							Val:   0,
							Left:  left,
							Right: right,
						})
					}
				}
			}
		}

		if len(trees) > 0 {
			completeBTrees[node] = trees
		}
	}

	return completeBTrees[n]
}
