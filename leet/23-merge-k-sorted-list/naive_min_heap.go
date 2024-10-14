package mergeksortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode

	hp := &MinHeap{
		arr: make([]*ListNode, 0, len(lists)),
	}

	for {
		for i := 0; i < len(lists); i++ {
			c := lists[i]
			if c == nil {
				lists = append(lists[:i], lists[i+1:]...)
				continue
			}

			hp.Insert(c)
			lists[i] = lists[i].Next
		}

		if len(lists) == 0 {
			break
		}
	}

	for {
		x := hp.ExtractMin()
		if x == nil {
			break
		}

		if head == nil {
			head = x
			tail = head
		} else {
			tail.Next = x
			tail = tail.Next
		}
	}

	return head
}

type MinHeap struct {
	arr []*ListNode
}

func (h *MinHeap) Insert(key *ListNode) {
	h.arr = append(h.arr, key)
	h.upHeap(len(h.arr) - 1)
}

func (h *MinHeap) ExtractMin() *ListNode {
	if len(h.arr) == 0 {
		return nil
	}
	root := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.downHeap(0)
	return root
}

func (h *MinHeap) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.arr[i].Val >= h.arr[parent].Val {
			break
		}
		h.arr[i], h.arr[parent] = h.arr[parent], h.arr[i]
		i = parent
	}
}

func (h *MinHeap) downHeap(i int) {
	n := len(h.arr)
	if n == 1 {
		return
	}

	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.arr[left].Val < h.arr[smallest].Val {
			smallest = left
		}
		if right < n && h.arr[right].Val < h.arr[smallest].Val {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		i = smallest
	}
}
