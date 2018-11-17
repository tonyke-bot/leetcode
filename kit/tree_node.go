package kit

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode a a tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNodeFromString builds a b-tree from given string splited with comma
func NewTreeNodeFromString(sequence string) *TreeNode {
	sequence = strings.Trim(sequence, " ")
	if len(sequence) == 0 || sequence == "null" {
		return nil
	}

	items := strings.Split(sequence, ",")
	rootVal, _ := strconv.Atoi(items[0])
	root := &TreeNode{Val: rootVal}

	queue := []*TreeNode{root}
	length := len(items)
	itemPos := 1

	for itemPos < length {
		root := queue[0]

		if itemPos < length && items[itemPos] != "null" {
			val, _ := strconv.Atoi(items[itemPos])
			root.Left = &TreeNode{Val: val}
			queue = append(queue, root.Left)
		}
		itemPos++

		if itemPos < length && items[itemPos] != "null" {
			val, _ := strconv.Atoi(items[itemPos])
			root.Right = &TreeNode{Val: val}
			queue = append(queue, root.Right)
		}
		itemPos++

		queue = queue[1:]
	}

	return root
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
