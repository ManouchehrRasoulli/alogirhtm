package removenth

import (
	"fmt"
	"testing"
)

func String(head *ListNode) string {
	if head == nil {
		return "nil"
	}

	return fmt.Sprintf("%d -> %s", head.Val, String(head.Next))
}

func Test_removeNthFromEnd(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	t.Log(String(h))

	x := removeNthFromEnd(h, 2)

	t.Log(String(x))
}

func Test_removeNthFromEnd2(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	t.Log(String(h))

	x := removeNthFromEnd_naive(h, 2)

	t.Log(String(x))
}

func Test_removeNthFromEnd3(t *testing.T) {
	h := &ListNode{
		Val: 1,
	}

	t.Log(String(h))

	x := removeNthFromEnd_naive(h, 1)

	t.Log(String(x))
}
