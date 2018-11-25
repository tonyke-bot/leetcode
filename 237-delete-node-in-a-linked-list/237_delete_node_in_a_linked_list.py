'''
Source: https://leetcode.com/problems/delete-node-in-a-linked-list
Test Case:
[4,5,1,9]
5

Solution:
    Remove forward all the successor item and unlink the last one
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def deleteNode(self, node):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
        if node is None:
            return

        while True:
            node.val = node.next.val

            if node.next.next is None:
                break
            node = node.next

        node.next = None


if __name__ == '__main__':
    solver = Solution()
