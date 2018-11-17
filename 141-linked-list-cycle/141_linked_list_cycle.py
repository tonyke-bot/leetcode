# Source: https://leetcode.com/problems/linked-list-cycle
class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        visited = set([])
        while head is not None:
            if head in visited:
                return True

            visited.add(head)
            head = head.next

        return False


# Test Case:
# []
# no cycle
