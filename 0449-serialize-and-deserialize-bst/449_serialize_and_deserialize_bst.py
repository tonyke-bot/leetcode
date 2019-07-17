'''
Source: https://leetcode.com/problems/serialize-and-deserialize-bst
Test Case:
[2,1,3]
'''


class TreeNode(object):
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

    def print_dfs(self, indent=""):
        if self is None:
            return

        print(indent + str(self.val))

        indent += "  "
        if self.left is None and self.right is not None:
            print(indent + "n")
            self.right.print_dfs(indent)
        else:
            if self.left:
                self.left.print_dfs(indent)
            if self.right:
                self.right.print_dfs(indent)


class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        if root is None:
            return ""

        result = []
        layer = [root]

        while len(layer):
            nextLayer = []
            for node in layer:
                result.append("n" if node is None else str(node.val))

                if node is not None:
                    nextLayer.append(node.left)
                    nextLayer.append(node.right)
            layer = nextLayer

        while result[-1] == "n":
            result.pop()

        return ",".join(result)

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        if len(data) == 0:
            return None

        vals = data.split(",")
        root = TreeNode(int(vals[0]))
        queue = [root]

        nodes = len(vals)
        i = 1
        while i < nodes:
            node = queue[0]

            if i < nodes and vals[i] != "n":
                node.left = TreeNode(int(vals[i]))
                queue.append(node.left)
            i += 1

            if i < nodes and vals[i] != "n":
                node.right = TreeNode(int(vals[i]))
                queue.append(node.right)
            i += 1

            queue = queue[1:]

        return root


# Your Codec object will be instantiated and called as such:
if __name__ == '__main__':
    codec = Codec()

    root1 = TreeNode(1, None, TreeNode(2, TreeNode(3), TreeNode(4)))
    root1_de = codec.deserialize(codec.serialize(None))
    root1_de.print_dfs()
