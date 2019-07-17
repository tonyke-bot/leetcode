package problem0096

/*
Source: https://leetcode.com/problems/unique-binary-search-trees
Test Case:
3

Solution:
	This problem can be solved with DP.
		d[n] =| Sum_(i=1, i->n)(d[i-1]*d[n-i])
			 =| 1, n <= 0,

	d[n] represents how many unique BST can be build from a n-length sorted array.

	First of all, we choose an element to be the root of the tree, and we can get
		the amount of nodes in the left subtree and in the right subtree.
	Then the amount of unique BST can be given the amount of unique BST for left
		subtree's length multiplies that of right subtree. Sth like combination.

	for example:
		for a 2-length array:
			1. root is the 0-th element: left subtree has 0 node, right subtree has 1 node
			2. root is the 0-th element: left subtree has 1 node, right subtree has 0 node
			Thus, count = 1*1 + 1*1 = 2

		for a 3-length array:
			1. root is the 0-th element: left subtree has 1 node, right subtree has 2 nodes
			2. root is the 1-st element: left subtree has 1 node, right subtree has 1 node
			3. root is the 0-th element: left subtree has 2 nodes, right subtree has 0 node
			Thus, count = 1*2 + 1*1 + 2*1 = 5
*/

func numTrees(n int) int {

	d := map[int]int{0: 1, 1: 1}

	for length := 2; length <= n; length++ {
		count := 0

		leftN := 0
		rightN := length - 1

		for leftN <= length-1 && rightN >= 0 {
			count += d[leftN] * d[rightN]

			leftN++
			rightN--
		}

		d[length] = count
	}

	return d[n]
}
