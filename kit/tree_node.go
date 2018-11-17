package kit

import "fmt"

// TreeNode a a tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PrintBFS prints the tree node BFS, empty solt is replace with nil
func (t *TreeNode) PrintBFS() {
	nextLayer := []*TreeNode{t}
	var currentLayer []*TreeNode

	for len(nextLayer) > 0 {
		currentLayer = nextLayer
		nextLayer = make([]*TreeNode, 0, len(currentLayer)*2)

		for _, node := range currentLayer {
			if node != nil {
				fmt.Print(node.Val, ',')
				nextLayer = append(nextLayer, node.Left)
				nextLayer = append(nextLayer, node.Right)
			} else {
				fmt.Print("nil", ',')
				nextLayer = append(nextLayer, nil)
				nextLayer = append(nextLayer, nil)
			}
		}
	}
}

// PrintDFS prints the tree node
func (t *TreeNode) PrintDFS() {
	var dfs func(string, *TreeNode)
	dfs = func(indent string, root *TreeNode) {
		if root == nil {
			return
		}

		fmt.Printf("%s%d\n", indent, root.Val)
		indent += "  "
		dfs(indent, root.Left)
		dfs(indent, root.Right)
	}

	dfs("", t)
}
