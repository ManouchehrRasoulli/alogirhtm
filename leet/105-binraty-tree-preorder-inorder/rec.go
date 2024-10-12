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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	head := TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	var inorderIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == head.Val {
			inorderIndex = i
			break
		}
	}

	head.Left = buildTree(preorder[1:inorderIndex+1], inorder[:inorderIndex])
	head.Right = buildTree(preorder[inorderIndex+1:], inorder[1+inorderIndex:])

	return &head
}
