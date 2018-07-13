from utils.list_node import ListNode, load_list, dump_list
from utils.input import input_int_list


class Solution:
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        root = ListNode(0)
        result = root

        carry = 0
        while l1 or l2 or carry:
            num = carry
            if l1:
                num += l1.val
                l1 = l1.next
            if l2:
                num += l2.val
                l2 = l2.next

            carry = num // 10
            result.val = num % 10

            if carry == 0 and l1 is None and l2 is None:
                return root
            else:
                result.next = ListNode(0)
                result = result.next


if __name__ == "__main__":
    solver = Solution()

    list_a = load_list(input_int_list("list a: "))
    list_b = load_list(input_int_list("list b: "))
    result = solver.addTwoNumbers(list_a, list_b)
    print(dump_list(result))
