package leetcode

// Source: https://leetcode.com/problems/merge-two-binary-trees

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	sum := 0
	if t1 != nil {
		sum += t1.Val
	}
	if t2 != nil {
		sum += t2.Val
	}

	if t1 == nil {
		return &TreeNode{
			Val:   sum,
			Left:  mergeTrees(nil, t2.Left),
			Right: mergeTrees(nil, t2.Right),
		}
	} else if t2 == nil {
		return &TreeNode{
			Val:   sum,
			Left:  mergeTrees(t1.Left, nil),
			Right: mergeTrees(t1.Right, nil),
		}
	}

	return &TreeNode{
		Val:   sum,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}

/*
Test Case:
[1,3,2,5]
[2,1,3,null,4,null,7]
*/
