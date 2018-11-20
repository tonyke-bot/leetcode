'''
Source: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii
Test Case:
{}
'''


class TreeLinkNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        self.next = None


class Solution:
    # @param root, a tree link node
    # @return nothing
    def connect(self, root):
        layer = []

        if root is not None:
            layer.append(root)

        while len(layer):
            nextLayer = []

            while None in layer:
                layer.remove(None)

            nodePos = 1
            while nodePos < len(layer):
                layer[nodePos-1].next = layer[nodePos]
                nodePos += 1

            for node in layer:
                if node.left is not None:
                    nextLayer.append(node.left)
                if node.right is not None:
                    nextLayer.append(node.right)

            layer = nextLayer
