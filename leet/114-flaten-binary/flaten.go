package flatenbinary

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return "nil"
	}

	return fmt.Sprintf("{val: %d, left: %s, right: %s}", t.Val, t.Left, t.Right)
}

func rightMost(node *TreeNode) *TreeNode {
	if node.Right == nil {
		return node
	}

	return rightMost(node.Right)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	if root.Left != nil {
		lm := rightMost(root.Left)
		lm.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	flatten(root.Right)
}
