package swapnodes

import "testing"

func Test_swapPairs(t *testing.T) {
	h := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	})

	t.Log(h)
}
