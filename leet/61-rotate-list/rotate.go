package rotatelist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l.Next == nil {
		return fmt.Sprintf("{val:%d}", l.Val)
	}
	return fmt.Sprintf("{val:%d} -> %s", l.Val, l.Next.String())
}

func (l *ListNode) Append(item *ListNode) {
	if l.Next == nil {
		l.Next = item
		return
	}

	l.Next.Append(item)
}

func rotateRight(head *ListNode, k int) *ListNode {
	/*
		solution
		1 - find the list-node length
		2 - cut the last k items if k > list-len then do nothings an re iterate with k - listLen
		3 - append head into tail of the new list
	*/

	if head == nil {
		return head
	}

	if k == 0 {
		return head
	}

	listLen := 1
	end := head
	for end.Next != nil {
		listLen++
		end = end.Next
	}

	if k >= listLen {
		return rotateRight(head, k%listLen)
	}

	end.Next = head

	newHead := head
	for i := 1; i < listLen-k; i++ {
		newHead = newHead.Next
	}
	head = newHead.Next
	newHead.Next = nil

	return head
}
