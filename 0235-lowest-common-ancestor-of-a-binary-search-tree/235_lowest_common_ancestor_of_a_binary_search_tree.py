'''
Source: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree
Test Case:
[6,2,8,0,4,7,9,null,null,3,5]
2
8
'''


# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right


class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        left, right = sorted([p.val, q.val])
        while root.val < left or right < root.val:
            root = root.right if left > root.val else root.left
        return root


if __name__ == '__main__':
    solver = Solution()
    node = TreeNode(6,
                    TreeNode(2,
                             TreeNode(0),
                             TreeNode(4,
                                      TreeNode(3),
                                      TreeNode(5))),
                    TreeNode(8,
                             TreeNode(7),
                             TreeNode(9)))

    result = solver.lowestCommonAncestor(node,
                                         node.left.right.left,
                                         node.left.right.right)
    print(result.val)
