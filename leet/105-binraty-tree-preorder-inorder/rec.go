package binratytreepreorderinorder

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func StringTree(root *TreeNode) string {
	if root == nil {
		return "nil"
	}

	return fmt.Sprintf("{val: %d, left: %s, right: %s}", root.Val, StringTree(root.Left), StringTree(root.Right))
}

func itemIndex(list []int, item int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return i
		}
	}

	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	head := TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	inorderIndex := itemIndex(inorder, val)
	inorderLeft := inorder[:inorderIndex]
	inorderRight := inorder[1+inorderIndex:]

	il := len(inorderLeft) + 1

	preOrderLeft := preorder[1:]
	preOrderRight := preorder[il:]

	head.Left = buildTree(preOrderLeft, inorderLeft)
	head.Right = buildTree(preOrderRight, inorderRight)

	return &head
}
