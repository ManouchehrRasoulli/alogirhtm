package removenth

type ListNode struct {
	Val  int
	Next *ListNode
}

func llen(head *ListNode) int {
	if head == nil {
		return 0
	}

	return 1 + llen(head.Next)
}

func removeNthFromEnd_naive(head *ListNode, n int) *ListNode {
	l := llen(head)

	if n > l {
		return head
	}

	p := head
	for i := 1; i < l-n; i++ {
		p = p.Next
	}

	if n == l {
		return head.Next
	}

	if p.Next.Next == nil {
		p.Next = nil
		return head
	}

	*p.Next = *p.Next.Next

	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headHolder := &ListNode{Next: head}
	faster := headHolder
	var i int
	for i = n; i >= 0; i-- {
		faster = faster.Next
	}

	if faster == nil && i >= 0 {
		return head
	}

	slower := headHolder
	for faster != nil {
		slower = slower.Next
		faster = faster.Next
	}

	slower.Next = slower.Next.Next
	return headHolder.Next
}
