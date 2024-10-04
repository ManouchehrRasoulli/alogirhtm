package rotatelist

import "testing"

func Test_RotateRight(t *testing.T) {
	nodes := listNodeHelper(1, 3)
	t.Log(nodes.String())

	rotated := rotateRight(nodes, 2000000000)
	t.Log(rotated.String())
}

func listNodeHelper(start, end int) *ListNode {
	head := &ListNode{
		Val:  start,
		Next: nil,
	}

	for i := start + 1; i <= end; i++ {
		head.Append(&ListNode{
			Val:  i,
			Next: nil,
		})
	}

	return head
}
