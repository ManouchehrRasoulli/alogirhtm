package swapnodes

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	f := head
	head = head.Next
	fp := f
	fp.Next = nil

	s := head
	sp := s
	head = head.Next
	sp.Next = nil

	i := 3
	for head != nil {
		if i%2 == 0 {
			sp.Next = head
			head = head.Next
			sp = sp.Next
			sp.Next = nil
		} else {
			fp.Next = head
			head = head.Next
			fp = fp.Next
			fp.Next = nil
		}
		i++
	}

	dummy := &ListNode{Next: nil}
	current := dummy

	for s != nil || f != nil {
		if s != nil {
			current.Next = s
			s = s.Next
			current = current.Next
		}

		if f != nil {
			current.Next = f
			f = f.Next
			current = current.Next
		}
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
