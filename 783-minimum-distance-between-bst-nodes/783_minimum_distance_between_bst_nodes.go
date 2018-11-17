package main

import "sort"

/*
Source: https://leetcode.com/problems/minimum-distance-between-bst-nodes
Test Case:
[4,2,6,1,3,null,null]
*/

type valueType []int

func (v valueType) Len() int           { return len(v) }
func (v valueType) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v valueType) Less(i, j int) bool { return v[i] < v[j] }

func minDiffInBST(root *TreeNode) int {
	nodes := make([]int, 0, 150)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		nodes = append(nodes, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	sort.Sort(valueType(nodes))

	minDiff := 1<<31 - 1
	for index := len(nodes) - 1; index >= 1; index-- {
		diff := nodes[index] - nodes[index-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
