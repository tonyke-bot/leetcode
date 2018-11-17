package leetcode

/*
Source: https://leetcode.com/problems/diameter-of-binary-tree
Test Case:
[1,2,3,4,5]
*/

func diameterOfBinaryTree(root *TreeNode) int {
	var maxLength = 0

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftLongest := dfs(root.Left)
		rightLongest := dfs(root.Right)

		path := leftLongest + rightLongest
		if path > maxLength {
			maxLength = path
		}

		if leftLongest > rightLongest {
			return leftLongest + 1
		}
		return rightLongest + 1
	}
	dfs(root)

	return maxLength
}
