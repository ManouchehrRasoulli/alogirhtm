package mergeksortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	hp := &MinHeap{
		arr: make([]*ListNode, 0, len(lists)),
	}

	for i := 0; i < len(lists); i++ {
		hp.Insert(lists[i])
	}

	dummy := &ListNode{}
	current := dummy

	for len(hp.arr) != 0 {
		smallest := hp.ExtractMin()
		if smallest == nil {
			break
		}
		current.Next = smallest
		current = current.Next
	}

	return dummy.Next
}

type MinHeap struct {
	arr []*ListNode
}

func (h *MinHeap) Insert(key *ListNode) {
	if key == nil {
		return
	}
	h.arr = append(h.arr, key)
	h.upHeap(len(h.arr) - 1)
}

func (h *MinHeap) ExtractMin() *ListNode {
	if len(h.arr) == 0 {
		return nil
	}
	root := h.arr[0]
	h.arr[0] = h.arr[0].Next
	if h.arr[0] == nil {
		h.arr[0] = h.arr[len(h.arr)-1]
		h.arr = h.arr[:len(h.arr)-1]
	}
	h.downHeap(0)
	root.Next = nil
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
