class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def load_list(elements):
    """list -> linked list"""
    root = ListNode(0)
    cursor = root
    for ele in elements:
        cursor.next = ListNode(ele)
        cursor = cursor.next
    return root.next


def dump_list(list_node):
    """linked list -> list"""
    elements = []
    while list_node:
        elements.append(list_node.val)
        list_node = list_node.next
    return elements
