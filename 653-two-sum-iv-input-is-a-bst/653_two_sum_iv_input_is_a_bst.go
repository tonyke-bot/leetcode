package leetcode

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/two-sum-iv-input-is-a-bst
Test Case:
[5,3,6,2,4,null,7]
9
*/

// TreeNode is a tree node
type TreeNode = kit.TreeNode

func findTarget(root *TreeNode, k int) bool {
	nums := []int{}
	count := map[int]int{}
	inOrderTraverse(root, &nums, &count)

	for _, num := range nums {
		count[num]--
		target := k - num
		times := count[target]
		if times > 0 {
			return true
		}
		count[num]++
	}

	return false
}

func inOrderTraverse(root *TreeNode, result *[]int, count *map[int]int) {
	if root != nil {
		inOrderTraverse(root.Left, result, count)
		*result = append(*result, root.Val)
		if _, found := (*count)[root.Val]; !found {
			(*count)[root.Val] = 0
		}
		(*count)[root.Val]++
		inOrderTraverse(root.Right, result, count)
	}
}
