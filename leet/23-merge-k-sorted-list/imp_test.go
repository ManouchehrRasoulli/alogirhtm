package mergeksortedlist

import (
	"fmt"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	x1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
		},
	}
	x2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	h := mergeKLists([]*ListNode{x1, x2})

	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
