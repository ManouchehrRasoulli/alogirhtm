package kd

import "fmt"

func PrintKDTree(root *node) {
	fmt.Println("Root")
	printNode(root, "", true)
}

func printNode(n *node, prefix string, isTail bool) {
	if n == nil {
		return
	}

	pointStr := fmt.Sprintf("%s (%.2f", n.tag, n.point[0])
	for i := 1; i < len(n.point); i++ {
		pointStr += fmt.Sprintf(", %.2f", n.point[i])
	}
	pointStr += fmt.Sprintf(" %d) [L%d]", n.repeat, n.level)

	connector := "├── "
	nextPrefix := prefix + "│   "
	if isTail {
		connector = "└── "
		nextPrefix = prefix + "    "
	}

	fmt.Println(prefix + connector + pointStr)

	children := []*node{}
	if n.left != nil {
		children = append(children, n.left)
	}
	if n.right != nil {
		children = append(children, n.right)
	}

	for i, child := range children {
		printNode(child, nextPrefix, i == len(children)-1)
	}
}
