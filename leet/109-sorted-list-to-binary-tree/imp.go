package sortedlisttobinarytree

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	t := TreeNode{
		Val: v,
	}

	return &t
}

func Height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(Height(node.Left), Height(node.Right))
}

func Insert(t *TreeNode, node *TreeNode) (head *TreeNode) {
	if t == nil {
		return node
	}

	if node.Val < t.Val {
		t.Left = Insert(t.Left, node)
	} else {
		t.Right = Insert(t.Right, node)
	}

	leftH := Height(t.Left)
	rightH := Height(t.Right)
	v := leftH - rightH

	if v < -1 {
		// rotate left
		parent := t
		head = parent.Right
		parent.Right = head.Left
		head.Left = parent
	} else if v > 1 {
		// rotate right
		parent := t
		head = parent.Left
		parent.Left = head.Right
		head.Right = parent
	} else {
		head = t
	}

	return head
}

func TreeToString(t *TreeNode) string {
	if t == nil {
		return "\"nil\""
	}

	return fmt.Sprintf("{\"val\": %d, \"left\": %s, \"right\": %s}", t.Val, TreeToString(t.Left), TreeToString(t.Right))
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	h := head
	t := NewTreeNode(h.Val)

	for h.Next != nil {
		h = h.Next
		t = Insert(t, NewTreeNode(h.Val))
	}

	return t
}
