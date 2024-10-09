package sortedlisttobinarytree

import "testing"

func Test_SortedListToBST(t *testing.T) {
	l := &ListNode{
		Val: 0,
		Next: &ListNode{
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
		},
	}

	tn := sortedListToBST(l)
	t.Log(TreeToString(tn))

	t.Log("done !!")
}
