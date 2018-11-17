'''
Source: https://leetcode.com/problems/n-ary-tree-level-order-traversal
Test Case:
{"$id":"1","children":[{"$id":"2","children":[{"$id":"5","children":[],"val":5},{"$id":"6","children":[],"val":6}],"val":3},{"$id":"3","children":[],"val":2},{"$id":"4","children":[],"val":4}],"val":1}
'''


# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children


class Solution(object):
    def levelOrder(self, root):
        """
        :type root: Node
        :rtype: List[List[int]]
        """
        result = []
        layer = []

        if root is not None:
            layer.append(root)

        while len(layer) > 0:
            vals = []
            nextLayer = []

            for node in layer:
                vals.append(node.val)
                for child in node.children:
                    if child is not None:
                        nextLayer.append(child)

            layer = nextLayer
            result.append(vals)

        return result


if __name__ == '__main__':
    solver = Solution()

    root = Node(1, [
        Node(3, [
            Node(5, []),
            Node(6, []),
        ]),
        Node(2, []),
        Node(4, []),
    ])

    print(solver.levelOrder(root))
